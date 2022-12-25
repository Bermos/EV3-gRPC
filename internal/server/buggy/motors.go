package buggy

import (
	"context"
	"github.com/Bermos/EV3-gRPC/internal/server/buggy/util"
	"github.com/ev3go/ev3dev"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var motorSet *MotorSet

type MotorsServerImpl struct {
	UnimplementedMotorsServer
}

func (m *MotorsServerImpl) On(_ context.Context, params *OnParams) (*Empty, error) {
	motorSet.on(parseSpeed(params))

	return &Empty{}, nil
}

func (m *MotorsServerImpl) OnForDegrees(_ context.Context, params *OnParams) (*Empty, error) {
	lSpeed, rSpeed := parseSpeed(params)
	motorSet.onForDegrees(lSpeed, rSpeed, int(params.GetDegrees()), params.GetBreak())

	return &Empty{}, nil
}

func (m *MotorsServerImpl) OnForRotations(_ context.Context, params *OnParams) (*Empty, error) {
	lSpeed, rSpeed := parseSpeed(params)
	motorSet.onForRotations(lSpeed, rSpeed, int(params.GetRotations()), params.GetBreak())

	return &Empty{}, nil
}

func (m *MotorsServerImpl) OnForSeconds(_ context.Context, params *OnParams) (*Empty, error) {
	lSpeed, rSpeed := parseSpeed(params)
	seconds := int(params.GetSeconds())
	if seconds < 0 {
		return &Empty{}, status.Error(codes.Canceled, "seconds cannot be smaller than 0")
	}

	motorSet.onForSeconds(lSpeed, rSpeed, seconds, params.GetBreak())

	return &Empty{}, nil
}

func (m *MotorsServerImpl) Left(_ context.Context, params *MotorParams) (*Empty, error) {
	speed, rampUp, rampDown, stopAction := parseMotorParams(params)

	motorSet.left.SetSpeedSetpoint(getNativeSpeed(speed))
	motorSet.left.SetRampUpSetpoint(time.Duration(rampUp) * time.Millisecond)
	motorSet.left.SetRampDownSetpoint(time.Duration(rampDown) * time.Millisecond)
	motorSet.left.SetStopAction(stopAction)

	return &Empty{}, nil
}

func (m *MotorsServerImpl) Right(_ context.Context, params *MotorParams) (*Empty, error) {
	speed, rampUp, rampDown, stopAction := parseMotorParams(params)

	motorSet.right.SetSpeedSetpoint(getNativeSpeed(speed))
	motorSet.right.SetRampUpSetpoint(time.Duration(rampUp) * time.Millisecond)
	motorSet.right.SetRampDownSetpoint(time.Duration(rampDown) * time.Millisecond)
	motorSet.right.SetStopAction(stopAction)

	return &Empty{}, nil
}

func parseMotorParams(params *MotorParams) (int, int, int, string) {
	speed := util.GetDefaultNumber(int(params.GetSpeed()), 10)
	rampUp := util.GetDefaultNumber(int(params.GetRampUp()), 5000)
	rampDown := util.GetDefaultNumber(int(params.GetRampDown()), 5000)
	stopAction := util.GetDefaultString(params.GetStop(), "hold")
	return speed, rampUp, rampDown, stopAction
}

func (m *MotorsServerImpl) Stop(_ context.Context, params *StopParams) (*Empty, error) {
	motorSet.stop(params.GetBreak())

	return &Empty{}, nil
}

func (m *MotorsServerImpl) WaitUntilNotMoving(_ context.Context, empty *Empty) (*Empty, error) {
	ev3dev.Wait(motorSet.left, ev3dev.Holding, 0, 0, false, 10*time.Second)
	ev3dev.Wait(motorSet.right, ev3dev.Holding, 0, 0, false, 10*time.Second)

	return &Empty{}, nil
}

func parseSpeed(params *OnParams) (int, int) {
	lSpeed, rSpeed := 10, 10
	if params.GetLSpeed() != 0 {
		lSpeed = int(params.GetLSpeed())
	}
	if params.GetRSpeed() != 0 {
		rSpeed = int(params.GetRSpeed())
	}
	return lSpeed, rSpeed
}
