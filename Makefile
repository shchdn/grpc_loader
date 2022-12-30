gen:
	protoc --go_out=./pkg/api --go-grpc_out=./pkg/api proto/uploader.proto

run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/*.go

run-tests:
	go test ./...