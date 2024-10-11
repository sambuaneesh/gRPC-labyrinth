package main

import (
	"context"
	"labyrinth/internal/game"
	labyrinth "labyrinth/protofiles"
	pb "labyrinth/protofiles"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GameStatus int

const (
	Running GameStatus = iota
	Win
	Death
)

type GameServer struct {
	Labyrinth *game.Labyrinth
	Player    *game.Player
	Status    GameStatus
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
		Status:    Running,
	}, nil
}

func (s *GameServer) GetLabyrinthInfo(ctx context.Context, req *pb.GetLabyrinthInfoRequest) (*pb.GetLabyrinthInfoResponse, error) {
	return &pb.GetLabyrinthInfoResponse{
		Width:  int32(s.Labyrinth.Width),
		Height: int32(s.Labyrinth.Height),
	}, nil
}

func (s *GameServer) GetPlayerStatus(ctx context.Context, req *pb.GetPlayerStatusRequest) (*pb.GetPlayerStatusResponse, error) {
	return &pb.GetPlayerStatusResponse{
		Score:           int32(s.Player.Score),
		Health:          int32(s.Player.Health),
		Position:        &pb.Position{X: int32(s.Player.Position.X), Y: int32(s.Player.Position.Y)},
		RemainingSpells: int32(s.Player.RemainingSpells),
	}, nil
}

func (s *GameServer) RegisterMove(ctx context.Context, req *pb.RegisterMoveRequest) (*pb.RegisterMoveResponse, error) {
	player := s.Player
	lab := s.Labyrinth
	initialPosition := player.Position

	if s.Status == Death || s.Status == Win {
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_FAILURE,
		}, nil
	}

	newPosition := player.Move(game.Direction(req.Direction))
	currentTile := lab.GetTile(newPosition)

	// victory condition
	if newPosition.X == lab.Width-1 && newPosition.Y == lab.Height-1 {
		s.Status = Win
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_VICTORY,
		}, nil
	}

	switch currentTile.Type {
	case game.Coin:
		player.CollectCoin()
		lab.SetTile(newPosition, game.Tile{Type: game.Empty})
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_SUCCESS,
		}, nil

	case game.Wall:
		player.LoseHealth()
		if player.Health == 0 {
			s.Status = Death
			return &pb.RegisterMoveResponse{
				Status: pb.MoveStatus_DEATH,
			}, nil
		}
		player.SetPosition(initialPosition)
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_FAILURE,
		}, nil

	case game.Empty:
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_SUCCESS,
		}, nil

	default:
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_FAILURE,
		}, nil
	}
}

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
	pb.RegisterLabyrinthGameServer(serverRegistrar, service)
	if err := serverRegistrar.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
