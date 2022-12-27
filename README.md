# EV3-gRPC

## Development

**Server**: 
 - Install dependencies: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
 - Generate Go code: `protoc --go_out=. --go-grpc_out=. ./protobuff/*.proto`
 - Compile server: `GOOS=linux GOARCH=arm GOARM=5 go build -o ev3api-server -ldflags="-s -w" cmd/main.go`

**Java client**:
 - Copy proto files: `cp protobuff/* clients/ev3api-java/src/main/proto`
 - Generate Java code: `[clients/ev3api-java] ./gradlew generateProto`
 - Create shadowed Jar: `[clients/ev3api-java] ./gradlew shadowJar`

**Python client**:
 - Install dependencies: `pip3 install grpcio-tools grpcio`
 - Generate Python code: `[clients/ev3api-python] python -m grpc_tools.protoc -I../../protobuff --python_out=. --pyi_out=. --grpc_python_out=. ../../protobuff/*.proto`
 - Build python wheel: `[clients/ev3api-python] python -m build`
