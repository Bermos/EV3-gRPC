package ev3

import (
	"context"
	"fmt"
	"github.com/ev3go/ev3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type LedServerImpl struct {
	UnimplementedLedServer
}

func (l *LedServerImpl) Flash(_ context.Context, led *EV3Led) (*Empty, error) {
	defer setColor("all", "off")

	color := "amber"
	if led.GetColor() != "" {
		color = led.GetColor()
	}

	err := setColor(led.GetSide(), color)
	if err != nil {
		return &Empty{}, status.Error(codes.Aborted, err.Error())
	}

	time.Sleep(500 * time.Millisecond)

	return &Empty{}, nil
}

func (l *LedServerImpl) Led(_ context.Context, led *EV3Led) (*Empty, error) {
	err := setColor(led.GetSide(), led.GetColor())
	if err != nil {
		return &Empty{}, status.Error(codes.Aborted, err.Error())
	}

	return &Empty{}, nil
}

func (l *LedServerImpl) LedOff(_ context.Context, _ *Empty) (*Empty, error) {
	err := setColor("all", "off")

	return &Empty{}, err
}

// ----------------------------- //

var colors = map[string]colorRG{
	"off":         {0, 0},
	"black":       {0, 0},
	"red":         {255, 0},
	"yellow":      {255, 26},
	"orange":      {255, 128},
	"amber":       {255, 255},
	"lime":        {128, 255},
	"green":       {0, 255},
	"dark_red":    {153, 0},
	"dar_orange":  {153, 76},
	"dark_yellow": {153, 153},
	"dark_lime":   {76, 153},
	"dark_green":  {0, 153},
}

type colorRG struct {
	Red   int
	Green int
}

func setColor(side, color string) error {
	color = strings.ToLower(color)
	c, ok := colors[color]
	if !ok {
		return fmt.Errorf("color '%s' not found", color)
	}

	switch side {
	case "left":
		ev3.RedLeft.SetBrightness(c.Red)
		ev3.GreenLeft.SetBrightness(c.Green)
	case "right":
		ev3.RedRight.SetBrightness(c.Red)
		ev3.GreenRight.SetBrightness(c.Green)
	case "all":
		fallthrough
	case "":
		ev3.RedLeft.SetBrightness(c.Red)
		ev3.GreenLeft.SetBrightness(c.Green)
		ev3.RedRight.SetBrightness(c.Red)
		ev3.GreenRight.SetBrightness(c.Green)
	default:
		return fmt.Errorf("side '%s' not found, allowed are 'left' and 'right'", side)
	}

	return nil
}
