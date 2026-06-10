package game

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/dakgg/go_grpc/server/internal/crypto"
	"github.com/dakgg/go_grpc/server/internal/model"
	pb "github.com/dakgg/go_grpc/server/proto"
)

type GameServer struct {
	pb.UnimplementedGameServiceServer
	DB *gorm.DB
}

func (s *GameServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	keyHash := crypto.HashPublicKey(req.Publickey)

	var user model.User
	result := s.DB.Where("public_key_hash = ?", keyHash).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		// 신규 유저 생성
		now := time.Now().UTC()
		user = model.User{
			PublicKeyHash: keyHash,
			CreatedAt:     now,
			LastLoginAt:   &now,
		}
		if err := s.DB.Create(&user).Error; err != nil {
			log.Printf("유저 생성 실패: %v", err)
			return &pb.LoginResponse{Success: false}, nil
		}
		log.Printf("신규 유저 생성: id=%d", user.ID)
	} else if result.Error != nil {
		log.Printf("DB 조회 실패: %v", result.Error)
		return &pb.LoginResponse{Success: false}, nil
	} else {
		// 기존 유저 last_login_at 업데이트
		now := time.Now().UTC()
		s.DB.Model(&user).Update("last_login_at", now)
		log.Printf("기존 유저 로그인: id=%d", user.ID)
	}

	return &pb.LoginResponse{
		Success:   true,
		GameState: "lobby",
	}, nil
}
