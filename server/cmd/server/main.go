package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	gameservice "github.com/dakgg/go_grpc/server/internal/game"
	pb "github.com/dakgg/go_grpc/server/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("포트 열기 실패: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, &gameservice.GameServer{})

	// grpcurl 같은 툴로 디버깅할 때 필요
	reflection.Register(s)

	log.Println("gRPC 서버 시작 :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("서버 실패: %v", err)
	}
}
