package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// loggingInterceptor は、リクエストが来たときにログを出力するインターセプター
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	log.Printf("Request - Method:%s\tBody:%v\n", info.FullMethod, req)
	h, err := handler(ctx, req)
	log.Printf("Response - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)
	return h, err
}
