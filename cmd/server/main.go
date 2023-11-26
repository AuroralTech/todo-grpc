package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/AuroralTech/todo-grpc/config"
	pb "github.com/AuroralTech/todo-grpc/pkg/grpc/generated"
	handler "github.com/AuroralTech/todo-grpc/pkg/handler/grpc"
	"github.com/AuroralTech/todo-grpc/pkg/infrastructure"
	"github.com/AuroralTech/todo-grpc/pkg/usecase"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. 5000番portのLisnterを作成
	port := 5000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	// 2. gRPCサーバーを作成
	s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))

	// 3. サーバーリフレクションの設定
	reflection.Register(s)

	// 4.依存関係の注入
	app := fx.New(
		fx.Provide(
			config.NewSQLConnection,
			infrastructure.NewTodoRepository,
			usecase.NewTodoUsecase,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, tu usecase.TodoUsecase) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						// 5.gRPCサーバーにTodoHandlerを登録
						pb.RegisterTodoServiceServer(s, handler.NewTodoHandler(tu))
						go func() {
							// 6.gRPCサーバーを起動
							log.Printf("start gRPC server port: %v", port)
							if err := s.Serve(listener); err != nil {
								log.Fatalf("failed to serve: %v", err)
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						// 7.Graceful shutdownされるようにする
						log.Println("stopping gRPC server...")
						s.GracefulStop()
						return nil
					},
				})

			},
		))

	app.Run()

}

// loggingInterceptor は、リクエストが来たときにログを出力するインターセプター
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	log.Printf("Request - Method:%s\tBody:%v\n", info.FullMethod, req)
	h, err := handler(ctx, req)
	log.Printf("Response - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)
	return h, err
}
