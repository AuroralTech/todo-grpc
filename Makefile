.PHONY: protoc-gen
protoc-gen:
	buf generate

.PHONY: protoc-lint
protoc-lint:
	buf lint

.PHONY: protoc-update
protoc-update:
	buf mod update

.PHONY: migrate
migrate:
	go run cmd/migration/main.go
