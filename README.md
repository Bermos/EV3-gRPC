# EV3-gRPC

## What and why

This project is a gRPC server for the Lego Mindstorms EV3. It allows you to control the EV3 from any programming
language that supports gRPC and has two client implementations for Java and Python respectively. This project was
developed as my project work (PA) for the [ZHAW](https://www.zhaw.ch/en/engineering/study/bachelors-degree-programmes/computer-science/)
together with another student. However, I decided to continue working on this project after the PA was finished.


## Development

### Prerequisites

Install GO: [Offical website](https://go.dev/dl/), version >= 1.19

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

#### Compile server for local use

This will build the server binary for your local machine. This is useful for testing the server locally.

 - Compile server: `go build -o server EV3-API/cmd`
