package ev3

import (
	"context"
	"github.com/ev3go/ev3dev"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const ps = ev3dev.PowerSupply("lego-ev3-battery")

type PowerServerImpl struct {
	UnimplementedPowerServer
}

func (p *PowerServerImpl) All(_ context.Context, _ *Empty) (*PowerInfo, error) {
	voltageMax, _ := ps.VoltageMax()
	voltageMin, _ := ps.VoltageMin()

	resp := &PowerInfo{}
	resp.Current, _ = ps.Current()
	resp.MaxVoltage = voltageMax * 1e-1 // It seems that our battery reports the wrong max and min voltages
	resp.MinVoltage = voltageMin * 1e-1 // It seems that our battery reports the wrong max and min voltages
	resp.Voltage, _ = ps.Voltage()
	resp.Technology, _ = ps.Technology()

	return resp, nil
}

func (p *PowerServerImpl) Current(_ context.Context, _ *Empty) (*PowerInfo, error) {
	current, err := ps.Current()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &PowerInfo{Current: current}, nil
}

func (p *PowerServerImpl) MaxVoltage(_ context.Context, _ *Empty) (*PowerInfo, error) {
	voltageMax, err := ps.VoltageMax()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &PowerInfo{MaxVoltage: voltageMax * 1e-1}, nil
}

func (p *PowerServerImpl) MinVoltage(_ context.Context, _ *Empty) (*PowerInfo, error) {
	voltageMin, err := ps.VoltageMin()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &PowerInfo{MinVoltage: voltageMin * 1e-1}, nil
}

func (p *PowerServerImpl) Technology(_ context.Context, _ *Empty) (*PowerInfo, error) {
	technology, err := ps.Technology()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &PowerInfo{Technology: technology}, nil
}

func (p *PowerServerImpl) Voltage(_ context.Context, _ *Empty) (*PowerInfo, error) {
	voltage, err := ps.Voltage()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &PowerInfo{Voltage: voltage}, nil
}
