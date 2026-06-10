package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dakgg/go_grpc/server/internal/db"
	gameservice "github.com/dakgg/go_grpc/server/internal/game"
	"github.com/dakgg/go_grpc/server/internal/interceptor"
	pb "github.com/dakgg/go_grpc/server/proto"
)

func main() {
	database := db.New(db.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "your_password",
		DBName:   "common",
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("포트 열기 실패: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Logging))
	pb.RegisterGameServiceServer(s, &gameservice.GameServer{DB: database})

	reflection.Register(s)

	log.Println("gRPC 서버 시작 :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("서버 실패: %v", err)
	}
}
