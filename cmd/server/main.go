package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/AuroralTech/todo-grpc/config"
	pb "github.com/AuroralTech/todo-grpc/pkg/generated"
	"github.com/AuroralTech/todo-grpc/pkg/infrastructure"
	handler "github.com/AuroralTech/todo-grpc/pkg/interfaces/grpc"
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

	// 4. サーバーリフレクションの設定
	reflection.Register(s)

	// 依存関係の注入
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			infrastructure.NewSQLConnection,
			infrastructure.NewTodoRepository,
			usecase.NewTodoUsecase,
			handler.NewTodoHandler,
		),
		fx.Invoke(
			func(tu *usecase.TodoUsecase) {
				pb.RegisterTodoServiceServer(s, handler.NewTodoHandler(*tu))
				go func() {
					log.Printf("start gRPC server port: %v", port)
					if err := s.Serve(listener); err != nil {
						log.Fatalf("failed to serve: %v", err)
					}
				}()
			},
		),
	)

	app.Run()

	// 6.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

// loggingInterceptor は、リクエストが来たときにログを出力するインターセプター
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	log.Printf("Request - Method:%s\tBody:%v\n", info.FullMethod, req)
	h, err := handler(ctx, req)
	log.Printf("Response - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)
	return h, err
}
