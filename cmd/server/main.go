package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/AuroralTech/todo-grpc/config"
	pb "github.com/AuroralTech/todo-grpc/pkg/generated"
	"github.com/AuroralTech/todo-grpc/pkg/infrastructure"
	handler "github.com/AuroralTech/todo-grpc/pkg/interfaces/grpc"
	"github.com/AuroralTech/todo-grpc/pkg/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. 5050番portのLisnterを作成
	port := 5050
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	// 依存関係の注入
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	db, err := infrastructure.NewSQLConnection(cfg.AppInfo.DatabaseURL)
	if err != nil {
		panic(err)
	}
	tr := infrastructure.NewTodoRepository(db)
	tu := usecase.NewTodoUsecase(tr)

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()

	// 3. gRPCサーバーにRegisterTodoServiceServerを登録
	pb.RegisterTodoServiceServer(s, handler.NewTodoHandler(*tu))

	// 4. サーバーリフレクションの設定
	reflection.Register(s)

	// 5. 作成したgRPCサーバーを、5000番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// 6.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
