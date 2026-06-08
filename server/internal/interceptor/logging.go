package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// Logging은 각 gRPC 요청의 처리 시간을 ms 단위로 로그에 남긴다.
func Logging(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()

	resp, err := handler(ctx, req)

	log.Printf("%s %dms", info.FullMethod, time.Since(start).Milliseconds())

	return resp, err
}
