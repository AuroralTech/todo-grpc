package grpc

import (
	"context"
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Authenticate(ctx context.Context) error {
	// 1.gRPCのメタデータからJWTトークンを取得
	md, _ := metadata.FromIncomingContext(ctx)
	auth := md.Get("authorization")
	if len(auth) == 0 {
		return status.Error(codes.Unauthenticated, "authorization token is not provided")
	}

	// 2. Bearerを除くトークンを取得
	idToken := strings.TrimPrefix(auth[0], "Bearer ")

	// 3.Firebase SDKの初期化
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("error initializing app: %v\n", err)
	}

	// 4.認証クライアントを取得
	client, err := app.Auth(ctx)
	if err != nil {
		return fmt.Errorf("error getting Auth client: %v\n", err)
	}

	// 5.トークンを検証
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return fmt.Errorf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)

	return nil
}
