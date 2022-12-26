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
	evt := getLastButtonEvent(true)
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
		log.Fatalf("failed to create button waiter: %v", err)
	}

	for e := range w.Events {

		lastButtonEvent = &Event{
			ButtonEvent: e,
			TimeStamp:   time.Now(),
		}
		log.Printf("DEBUG - %+v\n", e)
	}
}

// getLastButtonEvent gets last ev3dev.ButtonEvent
func getLastButtonEvent(clear bool) *Event {
	if lastButtonEvent == nil {
		return nil
	}

	btnEvt := *lastButtonEvent
	if clear {
		lastButtonEvent = nil
	}

	return &btnEvt
}
