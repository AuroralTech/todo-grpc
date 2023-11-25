# Dockerfile
FROM golang:1.21.4-alpine3.18 AS builder

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# ワーキングディレクトリの設定
WORKDIR /go/src/app

# ModuleモードをON
ENV GO111MODULE=on

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . .

# makeとprotocのインストール
RUN apk add --no-cache make protobuf protobuf-dev git

# bufのインストール
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# AIRのインストール
RUN go install github.com/cosmtrek/air@latest

# go.modを参照し、go.sumファイルの更新を行う
RUN go mod tidy

EXPOSE 5000

# localではホットリロードを有効にしたいのでairで起動する
CMD ["air", "-c", ".air.toml"]
