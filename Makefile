gen:
	protoc --go_out=./pkg/api --go-grpc_out=./pkg/api proto/uploader.proto

run-server:
	go run . server

run-client:
	go run . client

run-tests:
	go test ./...