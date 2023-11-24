.PHONY: protoc-gen
protoc-gen:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. pkg/proto/*.proto
