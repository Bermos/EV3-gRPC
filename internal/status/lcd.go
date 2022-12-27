package status

import (
	"fmt"
	"github.com/ev3go/ev3"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	loadedFont *truetype.Font
)

type Pixel struct {
	X     int
	Y     int
	Color color.Color
}

func init() {
	var err error
	loadedFont, err = loadFont()
	if err != nil {
		log.Printf("ERROR - can't loading font: %v", err)
		return
	}

	err = ev3.LCD.Init(true)
	if err != nil {
		log.Printf("ERROR - can't initialise framebuffer: %v", err)
		return
	}
}

// write the given text to the EV3 LCD (display) after clearing the screen
func write(textContent string) (err error) {
	bgColor := color.White
	bg := image.NewUniform(bgColor)
	draw.Draw(ev3.LCD, ev3.LCD.Bounds(), bg, image.Pt(0, 0), draw.Src)

	return fastWrite(textContent)
}

// fastWrite the given text to the EV3 LCD (display) without clearing the screen first
func fastWrite(textContent string) (err error) {
	fgColor := color.Black
	fontSize := float64(13)

	code := strings.Replace(textContent, "\t", "    ", -1) // convert tabs into spaces
	text := strings.Split(code, "\n")                      // split newlines into arrays

	fg := image.NewUniform(fgColor)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(loadedFont)
	c.SetFontSize(fontSize)
	c.SetClip(ev3.LCD.Bounds())
	c.SetDst(ev3.LCD)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	textXOffset := 0
	textYOffset := 0 + int(c.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, s := range text {
		_, err = c.DrawString(strings.Replace(s, "\r", "", -1), pt)
		if err != nil {
			return
		}
		pt.Y += c.PointToFixed(fontSize * 1.2)
	}

	return
}

func showSystemTTY(b bool) {
	tty := 5
	if b {
		tty = 2
	}
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo %s | sudo -S chvt %d", "maker", tty))

	if err := cmd.Start(); err != nil {
		log.Printf("ERROR - cmd.Start: %v", err)
	}
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() != 0 {
			log.Printf("ERROR - can't run chvt cmd. %v", err)
		}
	}
}

func loadFont() (f *truetype.Font, err error) {
	fontFile := "/usr/share/fonts/truetype/dejavu/DejaVuSansMono.ttf"
	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		return
	}
	f, err = freetype.ParseFont(fontBytes)
	if err != nil {
		return
	}
	return
}
