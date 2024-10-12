package main

import (
	"labyrinth/generated/go/labyrinth"
	"labyrinth/internal/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverRegistrar := grpc.NewServer()
	service, err := server.NewGameServer()
	if err != nil {
		log.Fatalf("failed to create game server: %v", err)
	}
	labyrinth.RegisterLabyrinthGameServer(serverRegistrar, service)
	if err := serverRegistrar.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
