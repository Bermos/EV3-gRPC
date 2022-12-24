package buggy

import (
	"github.com/ev3go/ev3dev"
	"log"
	"time"
)

func init() {
	left, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")
	if err != nil {
		log.Printf("ERROR - could not initialize left motor")
		return
	}

	right, err := ev3dev.TachoMotorFor("ev3-ports:outD", "lego-ev3-l-motor")
	if err != nil {
		log.Printf("ERROR - could not initialize right motor")
		return
	}

	motorSet = NewMotorSet(&motor{left}, &motor{right})
}

func getNativeSpeed(speedPercent int) int {
	return 1050 * speedPercent / 100
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type MotorSet struct {
	left  *motor
	right *motor
}

func NewMotorSet(left *motor, right *motor) *MotorSet {
	return &MotorSet{left: left, right: right}
}

func (m *MotorSet) on(lSpeed, rSpeed int) {
	m.left.SetSpeedSetpoint(getNativeSpeed(lSpeed))
	m.right.SetSpeedSetpoint(getNativeSpeed(rSpeed))

	m.left.Command("run-forever")
	m.right.Command("run-forever")
}

func (m *MotorSet) onForDegrees(lSpeedPct, rSpeedPct, degrees int, brake bool) {
	lSpeed := getNativeSpeed(lSpeedPct)
	rSpeed := getNativeSpeed(rSpeedPct)
	m.left.SetSpeedSetpoint(getNativeSpeed(lSpeedPct))
	m.right.SetSpeedSetpoint(getNativeSpeed(rSpeedPct))

	var lDegrees, rDegrees int

	if degrees == 0 || (lSpeed == 0 && rSpeed == 0) {
		lDegrees = 0
		rDegrees = 0
	} else if abs(lSpeed) > abs(rSpeed) {
		lDegrees = degrees
		rDegrees = abs(rSpeed/lSpeed) * degrees
	} else {
		lDegrees = abs(lSpeed/rSpeed) * degrees
		rDegrees = degrees
	}

	m.left.setRelPositionDegreesAndSpeedSP(lDegrees, lSpeed)
	m.right.setRelPositionDegreesAndSpeedSP(rDegrees, rSpeed)
	m.left.setBrake(brake)
	m.right.setBrake(brake)

	m.left.Command("run-to-rel-pos")
	m.right.Command("run-to-rel-pos")
}
func (m *MotorSet) onForRotations(lSpeedPct, rSpeedPct, degrees int, brake bool) {
	m.onForDegrees(lSpeedPct, rSpeedPct, degrees*360, brake)
}

func (m *MotorSet) onForSeconds(lSpeedPct, rSpeedPct, seconds int, brake bool) {
	m.left.SetSpeedSetpoint(getNativeSpeed(lSpeedPct))
	m.right.SetSpeedSetpoint(getNativeSpeed(rSpeedPct))

	m.left.SetTimeSetpoint(time.Second * time.Duration(seconds))
	m.right.SetTimeSetpoint(time.Second * time.Duration(seconds))

	m.left.setBrake(brake)
	m.right.setBrake(brake)

	m.left.Command("run-timed")
	m.right.Command("run-timed")
}

func (m *MotorSet) stop(brake bool) {
	m.left.setBrake(brake)
	m.right.setBrake(brake)

	m.left.Command("stop")
	m.right.Command("stop")
}

type motor struct {
	*ev3dev.TachoMotor
}

func (m *motor) setBrake(brake bool) {
	if brake {
		m.SetStopAction("hold")
	} else {
		m.SetStopAction("coast")
	}
}

func (m *motor) setRelPositionDegreesAndSpeedSP(degrees, speed int) {
	if speed < 0 {
		degrees = -degrees
	}
	speed = abs(speed)
	posDelta := degrees * m.CountPerRot() / 360

	m.SetPositionSetpoint(posDelta)
	m.SetSpeedSetpoint(speed)
}
