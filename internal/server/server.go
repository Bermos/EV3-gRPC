package server

import (
	"github.com/Bermos/EV3-gRPC/internal/server/buggy"
	"github.com/Bermos/EV3-gRPC/internal/server/ev3"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not listen on address %s: %v", address, err)
	}

	grpcServer := grpc.NewServer()

	ev3.RegisterButtonServer(grpcServer, &ev3.ButtonServerImpl{})
	ev3.RegisterLedServer(grpcServer, &ev3.LedServerImpl{})
	ev3.RegisterPowerServer(grpcServer, &ev3.PowerServerImpl{})
	ev3.RegisterSoundServer(grpcServer, &ev3.SoundServerImpl{})

	buggy.RegisterMotorsServer(grpcServer, &buggy.MotorsServerImpl{})
	buggy.RegisterSensorsServer(grpcServer, &buggy.SensorsServerImpl{})

	log.Printf("INFO - starting server on address %s", address)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
