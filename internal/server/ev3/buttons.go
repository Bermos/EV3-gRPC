package ev3

import (
	"context"
	"github.com/ev3go/ev3dev"
	"log"
	"time"
)

type ButtonServerImpl struct {
	UnimplementedButtonServer
}

func (b *ButtonServerImpl) Pressed(_ context.Context, _ *Empty) (*Buttons, error) {
	evt := GetLastButtonEvent(false)
	resp := &Buttons{
		Pressed: evt != nil && time.Now().Sub(evt.TimeStamp) < 3*time.Second,
	}

	return resp, nil
}

// -------------------------------------- //

type Event struct {
	ev3dev.ButtonEvent
	TimeStamp time.Time
}

var (
	lastButtonEvent *Event
)

func init() {
	go wait()
	log.Printf("INFO - button event loop started")
}

func wait() {
	w, err := ev3dev.NewButtonWaiter()
	if err != nil {
		log.Printf("ERROR - failed to create button waiter: %v", err)
		return
	}
	defer w.Close()

	for e := range w.Events {
		if e.Type == 1 && e.Value == 1 {
			lastButtonEvent = &Event{
				ButtonEvent: e,
				TimeStamp:   time.Now(),
			}
			log.Printf("DEBUG - %+v", lastButtonEvent)
		}
	}
}

// GetLastButtonEvent gets last ev3dev.ButtonEvent
func GetLastButtonEvent(clear bool) *Event {
	if lastButtonEvent == nil {
		return nil
	}

	btnEvt := *lastButtonEvent
	if clear {
		lastButtonEvent = nil
	}

	return &btnEvt
}
