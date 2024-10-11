package main

import (
	"context"
	"labyrinth/internal/game"
	labyrinth "labyrinth/protofiles"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GameServer struct {
	Labyrinth *game.Labyrinth
	Player    *game.Player
	labyrinth.UnimplementedLabyrinthGameServer
}

func NewGameServer() (*GameServer, error) {
	labyrinth, err := game.NewLabyrinth("./labyrinth.txt")
	if err != nil {
		return nil, err
	}
	return &GameServer{
		Labyrinth: labyrinth,
		Player:    game.NewPlayer(),
	}, nil
}

func (s *GameServer) GetLabyrinthInfo(ctx context.Context, req *labyrinth.GetLabyrinthInfoRequest) (*labyrinth.GetLabyrinthInfoResponse, error) {
	return &labyrinth.GetLabyrinthInfoResponse{
		Width:  int32(s.Labyrinth.Width),
		Height: int32(s.Labyrinth.Height),
	}, nil
}

// func (s *GameServer) GetPlayerStatus() (int, int, game.Position, int) {
// 	return s.Player.Score, s.Player.Health, s.Player.Position, s.Player.RemainingSpells
// }

// func (s *GameServer) RegisterMove(direction game.Direction) (game.MoveStatus, error) {
// 	player := s.Player
// 	labyrinth := s.Labyrinth
// 	initialPosition := player.Position

// 	newPosition := player.Move(direction)
// 	currentTile := labyrinth.GetTile(newPosition)

// 	// Check for victory condition
// 	if newPosition.X == labyrinth.Width-1 && newPosition.Y == labyrinth.Height-1 {
// 		return game.VICTORY, nil
// 	}

// 	switch currentTile.Type {
// 	case game.Coin:
// 		player.CollectCoin()
// 		labyrinth.SetTile(newPosition, game.Tile{Type: game.Empty})
// 		return game.SUCCESS, nil

// 	case game.Wall:
// 		player.LoseHealth()
// 		if player.Health == 0 {
// 			return game.DEATH, nil
// 		}
// 		player.SetPosition(initialPosition)
// 		return game.FAILURE, nil

// 	case game.Empty:
// 		return game.SUCCESS, nil

// 	default:
// 		return game.FAILURE, nil
// 	}
// }

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverRegistrar := grpc.NewServer()
	service, err := NewGameServer()
	if err != nil {
		log.Fatalf("failed to create game server: %v", err)
	}
	labyrinth.RegisterLabyrinthGameServer(serverRegistrar, service)
	if err := serverRegistrar.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
