package main

import (
	"EV3-gRPC/internal/server/buggy"
	"EV3-gRPC/internal/server/ev3"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

var address = flag.String("address", ":9000", "address to bind to, can also be just the port ':9000'")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("Could not listen on address %s: %v", ":9000", err)
	}

	s := grpc.NewServer()

	ev3.RegisterButtonServer(s, &ev3.ButtonServerImpl{})
	ev3.RegisterLedServer(s, &ev3.LedServerImpl{})
	ev3.RegisterPowerServer(s, &ev3.PowerServerImpl{})
	ev3.RegisterSoundServer(s, &ev3.SoundServerImpl{})

	buggy.RegisterMotorsServer(s, &buggy.MotorsServerImpl{})
	buggy.RegisterSensorsServer(s, &buggy.SensorsServerImpl{})

	log.Printf("INFO - starting server on address %s", *address)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
