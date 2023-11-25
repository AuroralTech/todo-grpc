.PHONY: protoc-gen
protoc-gen:
	protoc  -I ${GOPATH}/src/app -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate --go_out=. --go-grpc_out=require_unimplemented_servers=false:. pkg/proto/*.proto

.PHONY: migrate
migrate:
	go run cmd/migration/main.go
