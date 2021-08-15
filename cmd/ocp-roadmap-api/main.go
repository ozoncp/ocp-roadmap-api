package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozoncp/ocp-roadmap-api/internal/api"
	ocp_roadmap_api "github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
)

func main() {
	go runJSON()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ocp_roadmap_api.RegisterOcpRoadmapApiServer(s, api.NewRoadmapAPI())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
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

	log.Println("Server starting...")
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}
