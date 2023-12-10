package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// loggingInterceptor は、リクエストが来たときにログを出力するインターセプター
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	log.Printf("Request - Method:%s\tBody:%v\n", info.FullMethod, req)

	h, err := handler(ctx, req)
	// エラーが発生した場合はログに出力する
	if err != nil {
		st, ok := status.FromError(err)
		// 認証エラーはログに出力しない
		if ok && st.Code() != codes.Unauthenticated {
			log.Printf("Error - Method:%s\tDuration:%s\tMessage:%v\n", info.FullMethod, time.Since(start), err)
		}
	}
	return h, err
}
