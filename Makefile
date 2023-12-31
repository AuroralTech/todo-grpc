# protoからpbを生成する
.PHONY: protoc-gen
protoc-gen:
	buf generate

# protoからdocumentを生成する
.PHONY: protoc-doc-gen
protoc-doc-gen:
	protoc --doc_out=./doc --doc_opt=markdown,index.md ./pkg/proto/*.proto

# protoのlintを実行する
.PHONY: protoc-lint
protoc-lint:
	buf lint

# buf.lockファイルのアップデートをする
.PHONY: protoc-update
protoc-update:
	buf mod update

# DBのマイグレーションをする
.PHONY: migrate
migrate:
	go run cmd/migration/main.go

# ネットワークの作成(1回だけ実行する)
.PHONY: create-network
create-network:
	docker network create bff_grpc_network