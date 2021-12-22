package rainbow

import (
	"image/color"
	"time"
)

type Channel int

const (
	Red Channel = iota
	Green
	Blue
)

type Rainbow struct {
	color *color.RGBA
	speed uint8
	phase Channel
}

// redcolor   = color.RGBA{255, 0, 0, 255}
// greencolor = color.RGBA{0, 255, 0, 255}
// bluecolor  = color.RGBA{0, 0, 255, 255}

// Create a new rainbow, speed 1-7, start phase rainbow.Red, rainbow.Green, or rainbow.Blue
func Create(speed uint8, start Channel) *Rainbow {
	switch start {
	case Red:
		return &Rainbow{&color.RGBA{255, 0, 0, 255}, speed, Green}
	case Green:
		return &Rainbow{&color.RGBA{0, 255, 0, 255}, speed, Blue}
	case Blue:
		return &Rainbow{&color.RGBA{0, 0, 255, 255}, speed, Red}
	default:
		return &Rainbow{&color.RGBA{255, 0, 0, 255}, speed, Green}
	}
}

// Update the color of the rainbow
func (r *Rainbow) Next() color.Color {
	switch r.phase {
	case Red:
		r.color.R += r.getSpeed()
		r.color.B -= r.getSpeed()
	case Green:
		r.color.G += r.getSpeed()
		r.color.R -= r.getSpeed()
	case Blue:
		r.color.B += r.getSpeed()
		r.color.G -= r.getSpeed()
	}
	r.inspectChannel()
	return r.color
}

// Translate speed 1-7 into factors of 255 for even transitions
func (r *Rainbow) getSpeed() uint8 {
	switch r.speed {
	case 1:
		return 1
	case 2:
		return 3
	case 3:
		return 5
	case 4:
		return 15
	case 5:
		return 17
	case 6:
		return 51
	default:
		return 85
	}
}

// Check if the color has reached the end of the channel, and if so, move to the next channel
func (r *Rainbow) inspectChannel() {
	switch *r.color {
	// red
	case color.RGBA{255, 0, 0, 255}:
		r.phase = Green
	// green
	case color.RGBA{0, 255, 0, 255}:
		r.phase = Blue
	// blue
	case color.RGBA{0, 0, 255, 255}:
		r.phase = Red
	}
}

// Get the rainbows current color
func (r *Rainbow) Color() *color.RGBA {
	return r.color
}

type RainbowAsync struct {
	rainbow *Rainbow
	ticker  *time.Duration
	quit    chan bool
}

func CreateAsync(speed time.Duration, start Channel) *RainbowAsync {
	ra := &RainbowAsync{
		Create(1, start),
		&speed,
		make(chan bool),
	}
	ra.async()
	return ra
}

func (ra *RainbowAsync) async() {
	go func() {
		for {
			select {
			case <-ra.quit:
				return
			default:
				time.Sleep(*ra.ticker)
				ra.rainbow.Next()
			}
		}
	}()
}

func (ra *RainbowAsync) Color() *color.RGBA {
	return ra.rainbow.Color()
}

func (ra *RainbowAsync) Close() {
	ra.quit <- true
}
