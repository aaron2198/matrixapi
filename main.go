package main

import (
	"flag"
	"os"
	"time"

	rgbmatrix "github.com/aaron2198/go-rpi-rgb-led-matrix"
	"github.com/aaron2198/go-rpi-rgb-led-matrix/display"
	"github.com/aaron2198/matrixapi/rainbow"
)

var (
	rows                     = flag.Int("led-rows", 32, "number of rows supported")
	cols                     = flag.Int("led-cols", 64, "number of columns supported")
	parallel                 = flag.Int("led-parallel", 1, "number of daisy-chained panels")
	chain                    = flag.Int("led-chain", 1, "number of displays daisy-chained")
	brightness               = flag.Int("brightness", 100, "brightness (0-100)")
	hardware_mapping         = flag.String("led-gpio-mapping", "regular", "Name of GPIO mapping used.")
	show_refresh             = flag.Bool("led-show-refresh", false, "Show refresh rate.")
	inverse_colors           = flag.Bool("led-inverse", false, "Switch if your matrix has inverse colors on.")
	disable_hardware_pulsing = flag.Bool("led-no-hardware-pulse", false, "Don't use hardware pin-pulse generation.")
)

func main() {
	config := &rgbmatrix.DefaultConfig
	config.Rows = *rows
	config.Cols = *cols
	config.Parallel = *parallel
	config.ChainLength = *chain
	config.Brightness = *brightness
	config.HardwareMapping = *hardware_mapping
	config.ShowRefreshRate = *show_refresh
	config.InverseColors = *inverse_colors
	config.DisableHardwarePulsing = *disable_hardware_pulsing

	// Create display configuration based on the C API.
	m, err := rgbmatrix.NewRGBLedMatrix(config)
	fatal(err)
	// Create a defualt window that uses the full display.
	def := display.CreateWindow(64, 32)
	// Create a new display with the default window.
	D := display.CreateDisplay(m, def, time.Millisecond*17)
	D.EnableDebugging(os.Stdout)
	// Create a new rainbow effect against a *color.RGBA.
	r := rainbow.CreateAsync(time.Millisecond*50, rainbow.Red)
	// Create a circle and give it the rainbow effect.
	circle := display.CreateCircle(32, 16, 5, r.Color())
	// Create a point bouncer
	bpanimation := display.CreateBouncePoint(64, 32, 5)
	circle.SetAnimation(bpanimation)
	// Add one circle to the default window
	D.Windows["default"].AddElement(circle)
	time.Sleep(time.Second * 10)
	// Create a second RGB circle and add it to the default window.
	r2 := rainbow.CreateAsync(time.Millisecond*50, rainbow.Green)
	circle2 := display.CreateCircle(32, 16, 5, r2.Color())
	D.Windows["default"].AddElement(circle2)
	bpanimation2 := display.CreateBouncePoint(64, 32, 5)
	circle2.SetAnimation(bpanimation2)
	time.Sleep(time.Second * 10)
	// Pause the display for a couple seconds.
	// D.ToState <- display.Stopped
	// time.Sleep(time.Second * 2)

	// Create a second window, add it to the displays window list.
	D.AddWindow("window2")
	circle3 := display.CreateCircle(32, 16, 3, rainbow.CreateAsync(time.Millisecond*50, rainbow.Green).Color())
	window2animation := display.CreateBouncePoint(64, 32, 3)
	circle3.SetAnimation(window2animation)
	D.Windows["window2"].AddElement(circle3)
	// Bring the second window to the front.
	D.Foreground("window2")
	D.ToState <- display.Running
	time.Sleep(time.Second * 10)
	// Bring the default window to the front.
	D.Foreground("default")
	time.Sleep(time.Second * 10)
	D.ToState <- display.Killed
}

func init() {
	flag.Parse()
}

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}

// type Animation struct {
// 	ctx      *gg.Context
// 	position image.Point
// 	dir      image.Point
// 	stroke   int
// 	rb       *rainbow.Rainbow
// 	State    rgbmatrix.State
// }

// func NewAnimation(sz image.Point) *Animation {
// 	rainbow := rainbow.Create(3, rainbow.Blue)
// 	return &Animation{
// 		ctx:    gg.NewContext(sz.X, sz.Y),
// 		dir:    image.Point{1, 1},
// 		stroke: 5,
// 		rb:     rainbow,
// 	}
// }

// func (a *Animation) Next() (image.Image, <-chan time.Time, error) {
// 	defer a.updatePosition()

// 	a.ctx.SetColor(color.Black)
// 	a.ctx.Clear()

// 	a.ctx.DrawCircle(float64(a.position.X), float64(a.position.Y), float64(a.stroke))
// 	a.ctx.SetColor(a.rb.Next())
// 	a.ctx.Fill()
// 	return a.ctx.Image(), time.After(time.Millisecond * 50), nil
// }

// func (a *Animation) SetState(s rgbmatrix.State) {
// 	a.State = s
// }

// func (a *Animation) updatePosition() {
// 	a.position.X += 1 * a.dir.X
// 	a.position.Y += 1 * a.dir.Y

// 	if a.position.Y+a.stroke > a.ctx.Height() {
// 		a.dir.Y = -1
// 	} else if a.position.Y-a.stroke < 0 {
// 		a.dir.Y = 1
// 	}

// 	if a.position.X+a.stroke > a.ctx.Width() {
// 		a.dir.X = -1
// 	} else if a.position.X-a.stroke < 0 {
// 		a.dir.X = 1
// 	}
// }

// type AnimationHandler struct {
// 	animation *Animation
// 	to        chan rgbmatrix.State
// }

// func (ah *AnimationHandler) handle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "toggling animation\n")

// 	switch ah.animation.State {
// 	case rgbmatrix.Stopped:
// 		fmt.Fprint(w, "starting animation\n")
// 		ah.to <- rgbmatrix.Running
// 	case rgbmatrix.Running:
// 		fmt.Fprint(w, "stopping animation\n")
// 		ah.to <- rgbmatrix.Stopped
// 	}

// }
