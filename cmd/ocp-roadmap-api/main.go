package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-roadmap-api/internal/api"
	db_connection "github.com/ozoncp/ocp-roadmap-api/internal/db-connection"
	"github.com/ozoncp/ocp-roadmap-api/internal/repo"
	ocp_roadmap_api "github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn := db_connection.Connection(ctx)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Info().Msgf("error while close connection DB, err: %s", err.Error())
		}
	}()

	srv := runJSON(ctx)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()

	gSrv, listen := runGRPC(conn)
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
}

func runGRPC(conn *sqlx.DB) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	repository := repo.NewRepository(conn)
	ocp_roadmap_api.RegisterOcpRoadmapApiServer(s, api.NewRoadmapAPI(repository))

	return s, listen
}

func runJSON(ctx context.Context) *http.Server {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := ocp_roadmap_api.RegisterOcpRoadmapApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	return srv
}
