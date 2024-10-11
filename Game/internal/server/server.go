package main

import (
	"context"
	"fmt"
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

	player := game.NewPlayer()

	GameServer := &GameServer{
		Labyrinth: labyrinth,
		Player:    player,
		Status:    Running,
	}
	GameServer.PrintLabyrinth()
	return GameServer, nil
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

	if s.Status == Death {
		s.PrintLabyrinth()
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_DEATH,
		}, nil
	} else if s.Status == Win {
		s.PrintLabyrinth()
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_VICTORY,
		}, nil
	}

	newPosition := player.Move(game.Direction(req.Direction))
	currentTile := lab.GetTile(newPosition)

	// victory condition
	if newPosition.X == lab.Width-1 && newPosition.Y == lab.Height-1 {
		s.Status = Win
		s.PrintLabyrinth()
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_VICTORY,
		}, nil
	}

	switch currentTile.Type {
	case game.Coin:
		player.CollectCoin()
		lab.SetTile(newPosition, game.Tile{Type: game.Empty})
		s.PrintLabyrinth()
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_SUCCESS,
		}, nil

	case game.Wall:
		player.LoseHealth()
		player.SetPosition(initialPosition)
		if player.Health == 0 {
			s.Status = Death
			s.PrintLabyrinth()
			return &pb.RegisterMoveResponse{
				Status: pb.MoveStatus_DEATH,
			}, nil
		}
		s.PrintLabyrinth()
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_FAILURE,
		}, nil

	case game.Empty:
		s.PrintLabyrinth()
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_SUCCESS,
		}, nil

	default:
		s.PrintLabyrinth()
		return &pb.RegisterMoveResponse{
			Status: pb.MoveStatus_FAILURE,
		}, nil
	}
}

func (s *GameServer) Revelio(req *pb.RevelioRequest, stream pb.LabyrinthGame_RevelioServer) error {
	targetX := int(req.Target.X)
	targetY := int(req.Target.Y)
	tileType := game.TileType(req.TileType)

	if s.Player.RemainingSpells == 0 {
		return fmt.Errorf("no remaining spells")
	}

	s.Player.RemainingSpells--

	// Check if the target position is valid
	if !s.Labyrinth.IsValidPosition(game.Position{X: targetX, Y: targetY}) {
		return fmt.Errorf("invalid target position")
	}

	// checking the target tile type
	if s.Labyrinth.GetTile(game.Position{X: targetX, Y: targetY}).Type == tileType {
		if err := stream.Send(&pb.RevelioResponse{
			Position: &pb.Position{X: int32(targetX), Y: int32(targetY)},
		}); err != nil {
			return err
		}
	}

	// defining for surrounding tiles
	directions := []game.Position{
		{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1},
		{X: -1, Y: 0}, {X: 1, Y: 0},
		{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
	}

	for _, dir := range directions {
		newX, newY := targetX+dir.X, targetY+dir.Y
		pos := game.Position{X: newX, Y: newY}

		if s.Labyrinth.IsValidPosition(pos) && s.Labyrinth.GetTile(pos).Type == tileType {
			if err := stream.Send(&pb.RevelioResponse{
				Position: &pb.Position{X: int32(newX), Y: int32(newY)},
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *GameServer) PrintLabyrinth() {
	fmt.Print("\n+")
	for i := 0; i < s.Labyrinth.Width; i++ {
		fmt.Print("--+")
	}
	fmt.Println()

	for y, row := range s.Labyrinth.Tiles {
		fmt.Print("|")
		for x, tile := range row {
			if x == s.Player.Position.X && y == s.Player.Position.Y {
				fmt.Print("🧙")
			} else {
				switch tile.Type {
				case game.Empty:
					fmt.Print("  ")
				case game.Wall:
					fmt.Print("🧱")
				case game.Coin:
					fmt.Print("💰")
				default:
					fmt.Print("❓")
				}
			}
			fmt.Print("|")
		}
		fmt.Println()

		fmt.Print("+")
		for i := 0; i < s.Labyrinth.Width; i++ {
			fmt.Print("--+")
		}
		fmt.Println()
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
