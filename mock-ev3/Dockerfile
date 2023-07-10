FROM golang:1.19-alpine AS go-builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd
COPY internal internal
COPY protobuff protobuff

RUN apk add protoc \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
    && protoc --go_out=. --go-grpc_out=. ./protobuff/*.proto
RUN GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

FROM debian

WORKDIR /app
COPY --from=go-builder /build/server ./server
RUN ls -ahl

ENTRYPOINT ["./server"]
