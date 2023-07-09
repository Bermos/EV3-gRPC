package buggy

import (
	"context"
	"github.com/ev3go/ev3dev"
	"log"
	"strconv"
)

type SensorsServerImpl struct {
	UnimplementedSensorsServer
}

func (s *SensorsServerImpl) Gyro(_ context.Context, _ *Empty) (*SensorResult, error) {
	Gyro.SetMode("GYRO-ANG")
	resp := &SensorResult{
		StrValue: Gyro.GetStringValue(),
		NumValue: int32(Gyro.GetNumValue() - gyroOffset),
	}

	return resp, nil
}

func (s *SensorsServerImpl) GyroReset(_ context.Context, _ *Empty) (*Empty, error) {
	gyroOffset = Gyro.GetNumValue()

	return &Empty{}, nil
}

func (s *SensorsServerImpl) Sonic(_ context.Context, _ *Empty) (*SensorResult, error) {
	Sonic.SetMode("US-DIST-CM")
	resp := &SensorResult{
		StrValue: Sonic.GetStringValue(),
		NumValue: int32(Sonic.GetNumValue()),
	}

	return resp, nil
}

// ----------------------------------------- //

var (
	Gyro       *sensor
	Sonic      *sensor
	gyroOffset = 0
)

type sensor struct {
	*ev3dev.Sensor
}

// GetStringValue returns the sensor value as a string
func (s *sensor) GetStringValue() string {
	v, err := s.Value(0)
	if err != nil {
		log.Printf("ERROR - sensor read error: %v", err)
		return ""
	}

	return v
}

// GetNumValue returns the sensor value as an int
func (s *sensor) GetNumValue() int {
	f, err := strconv.Atoi(s.GetStringValue())
	if err != nil {
		return 0
	}

	return f
}

func init() {
	gyro, err := ev3dev.SensorFor("ev3-ports:in4", "lego-ev3-gyro")
	if err != nil {
		log.Printf("ERROR - could not load gyro sensor: %v", err)
	} else {
		Gyro = &sensor{gyro}
		log.Println("INFO - Gyro sensor loaded")
	}

	sonic, err := ev3dev.SensorFor("ev3-ports:in1", "lego-ev3-us")
	if err != nil {
		log.Printf("ERROR - could not load sonic sensor: %v", err)
	} else {
		Sonic = &sensor{sonic}
		log.Println("INFO - Sonic sensor loaded")
	}

}
