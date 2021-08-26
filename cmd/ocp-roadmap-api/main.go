package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-roadmap-api/internal/api"
	cnfg "github.com/ozoncp/ocp-roadmap-api/internal/config"
	db_connection "github.com/ozoncp/ocp-roadmap-api/internal/db-connection"
	"github.com/ozoncp/ocp-roadmap-api/internal/kafka"
	"github.com/ozoncp/ocp-roadmap-api/internal/metric"
	"github.com/ozoncp/ocp-roadmap-api/internal/repo"
	"github.com/ozoncp/ocp-roadmap-api/internal/tracing"
	ocp_roadmap_api "github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var config *cnfg.Config

func main() {
	config = cnfg.InitConfig(cnfg.CONFIG_NAME)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tracer := tracing.InitTracing()

	conn := db_connection.Connection(ctx)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Info().Msgf("error while close connection DB, err: %s", err.Error())
		}
	}()

	metricServer := metricsSRV()
	go func() {
		if err := metricServer.ListenAndServe(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()

	srv := runJSON(ctx)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()
	prod := kafkaProducer()
	metric.InitMetrics()

	gSrv, listen := runGRPC(conn, prod)
	go func() {
		if err := gSrv.Serve(listen); err != nil {
			log.Fatal().Msgf("failed to serve: %v", err)
		}
	}()

	// Stop by signal
	<-c
	gSrv.GracefulStop()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msg(err.Error())
	}

	if err := tracer.Close(); err != nil {
		log.Err(err).Msg(err.Error())
	}

	if err := metricServer.Shutdown(ctx); err != nil {
		log.Err(err).Msg(err.Error())
	}
}

func kafkaProducer() kafka.Producer {
	prod := kafka.InitProducer()

	log.Info().Msg("Kafka message broker started and init")
	return prod
}

func runGRPC(conn *sqlx.DB, producer kafka.Producer) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPC.Port))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	repository := repo.NewRepository(conn)
	ocp_roadmap_api.RegisterOcpRoadmapApiServer(s, api.NewRoadmapAPI(repository, producer))

	return s, listen
}

func runJSON(ctx context.Context) *http.Server {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := ocp_roadmap_api.RegisterOcpRoadmapApiHandlerFromEndpoint(ctx, mux, config.REST.Endpoint, opts)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.REST.Port),
		Handler: mux,
	}
	return srv
}

func metricsSRV() *http.Server {
	sm := http.NewServeMux()
	sm.Handle(config.Metrics.Handle, promhttp.Handler())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Metrics.Port),
		Handler: sm,
	}

	return srv
}
