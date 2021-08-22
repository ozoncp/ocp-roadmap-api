package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jackc/pgx/v4"
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

	conn := db_connection.NewPGConnection(ctx)
	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Info().Msgf("error while close connection DB, err: %s", err.Error())
		}
	}()

	go runJSON()
	if err := runGRPC(conn); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func runGRPC(conn *pgx.Conn) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	repository := repo.NewRepository(conn)
	ocp_roadmap_api.RegisterOcpRoadmapApiServer(s, api.NewRoadmapAPI(repository))

	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := ocp_roadmap_api.RegisterOcpRoadmapApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	log.Info().Msg("Server starting...")
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}
