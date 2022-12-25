# EV3-gRPC

## Development

**Generate Go code**: `protoc --go_out=. --go-grpc_out=. ./protobuff/*.proto`


**Compile server:** `GOOS=linux GOARCH=arm GOARM=5 go build -o ev3api-server -ldflags="-s -w" cmd/main.go`
