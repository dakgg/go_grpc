package game

import (
	"context"
	"log"

	pb "github.com/dakgg/go_grpc/server/proto"
)

type GameServer struct {
	pb.UnimplementedGameServiceServer
}

func (s *GameServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Login 요청: publickey=%s", req.Publickey)

	println("Login")

	return &pb.LoginResponse{
		Success:   true,
		GameState: "lobby",
	}, nil
}
