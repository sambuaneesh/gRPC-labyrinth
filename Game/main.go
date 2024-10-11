package main

import (
	"context"
	labyrinth "labyrinth/protofiles"
	"log"
	"net"

	"google.golang.org/grpc"
)

type labyrinthServer struct {
	labyrinth.UnimplementedLabyrinthGameServer
}

func (s *labyrinthServer) GetLabyrinthInfo(ctx context.Context, req *labyrinth.GetLabyrinthInfoRequest) (*labyrinth.GetLabyrinthInfoResponse, error) {
	return &labyrinth.GetLabyrinthInfoResponse{
		Width:  10,
		Height: 10,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serviceRegistrar := grpc.NewServer()
	service := &labyrinthServer{}

	labyrinth.RegisterLabyrinthGameServer(serviceRegistrar, service)

	if err := serviceRegistrar.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
