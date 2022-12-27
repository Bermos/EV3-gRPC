package status

import (
	"fmt"
	"github.com/Bermos/EV3-gRPC/internal/server/buggy"
	"github.com/Bermos/EV3-gRPC/internal/server/ev3"
	"github.com/Bermos/EV3-gRPC/internal/util"
	"log"
	"net"
	"strings"
	"time"
)

var (
	updates = []update{
		{"IP", getIP},
		{"US", getUS},
		{"Gyro", getGyro},
	}

	lastDText string
	interval  time.Duration
)

type update struct {
	Name string
	Fn   func() string
}

// Start the update loop with the given interval between updates.
// The update loop runs in a separate go routine (almost like a thread).
// The duration between updates can go to 0 if the update takes longer than the interval
func Start(i time.Duration) {
	interval = i
	go startLoop()
}

func startLoop() {
	showSystemTTY(false)      // Hide system screen
	defer showSystemTTY(true) // Show system screen on exit

	for true {
		start := time.Now()
		displayStatus()

		evt := ev3.GetLastButtonEvent(false)
		if evt != nil && time.Now().Sub(evt.TimeStamp) < interval {
			break
		}

		duration := time.Now().Sub(start)
		// log.Printf("DEBUG - displayStatus: duration %v", duration)
		time.Sleep(interval - duration)
	}
}

func displayStatus() {
	dLines := make([]string, len(updates))
	for i, u := range updates {
		dLines[i] = fmt.Sprintf("%s: %s", u.Name, u.Fn())
	}

	dText := strings.Join(dLines, "\n")
	if dText != lastDText {
		err := write(dText)
		if err != nil {
			log.Println(err)
		}
		lastDText = dText
	} else {
		err := fastWrite(dText)
		if err != nil {
			log.Println(err)
		}
	}
}

func getIP() string {
	ifIdx := util.GetWlanInterfaceIndex()

	if intf, err := net.InterfaceByIndex(ifIdx); err == nil {
		if addrs, err := intf.Addrs(); err == nil {
			return addrs[0].String()
		}
	}

	return "not connected"
}

func getUS() string {
	if buggy.Sonic == nil {
		return "no sonic sensor"
	}

	return buggy.Sonic.GetStringValue()
}

func getGyro() string {
	if buggy.Gyro == nil {
		return "no gyro sensor"
	}

	return buggy.Gyro.GetStringValue()
}
