package ev3

import (
	"context"
	"fmt"
	"github.com/ev3go/ev3dev"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os/exec"
	"time"
)

const (
	soundPath  = "/dev/input/by-path/platform-sound-event"
	aplayPath  = "/usr/bin/aplay"
	espeakPath = "/usr/bin/espeak"
)

var speaker = ev3dev.NewSpeaker(soundPath)

func init() {
	err := speaker.Init()
	if err != nil {
		log.Printf("ERROR - failed to initialize speaker: %v", err)
		return
	}

	log.Printf("INFO - initialzed speaker")
}

type SoundServerImpl struct {
	UnimplementedSoundServer
}

func (s SoundServerImpl) Beep(_ context.Context, _ *Empty) (*Empty, error) {
	if err := playTone(&Tone{Frequency: 440, DurationMs: 200}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &Empty{}, nil
}

func (s SoundServerImpl) PlayTone(_ context.Context, tone *Tone) (*Empty, error) {
	if tone.GetFrequency() == 0 || tone.GetDurationMs() == 0 {
		return nil, status.Error(codes.Canceled, "missing content parameter")
	}

	if err := playTone(tone); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &Empty{}, nil
}

func (s SoundServerImpl) Speak(_ context.Context, text *Text) (*Empty, error) {
	content := text.GetContent()
	if content == "" {
		return nil, status.Error(codes.Canceled, "missing content parameter")
	}

	// FIXME: currently a ' will cause an error since it isn't escaped -> escape it
	cmd := fmt.Sprintf("%s --stdout -a 200 -s 130 '%s' | %s -q", espeakPath, content, aplayPath)
	if out, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		return nil, status.Errorf(codes.Internal, "%s - %v", out, err)
	}
	return &Empty{}, nil
}

// ----------------------------- //

func playTone(tone *Tone) (err error) {
	defer speaker.Tone(0)

	if err = speaker.Tone(tone.GetFrequency()); err != nil {
		return
	}
	time.Sleep(time.Duration(tone.GetDurationMs()) * time.Millisecond)
	return
}
