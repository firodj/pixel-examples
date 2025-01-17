package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
)

const w, h = float64(1024), float64(512)

var speed = float64(200)

var stars [1024]*star

func init() {
	rand.Seed(4)

	for i := 0; i < len(stars); i++ {
		stars[i] = newStar()
	}
}

type star struct {
	pixel.Vec
	Z float64
	P float64
	C color.RGBA
}

func newStar() *star {
	return &star{
		pixel.V(random(-w, w), random(-h, h)),
		random(0, w), 0, Colors[rand.Intn(len(Colors))],
	}
}

func (s *star) update(d float64) {
	s.P = s.Z
	s.Z -= d * speed

	if s.Z < 0 {
		s.X = random(-w, w)
		s.Y = random(-h, h)
		s.Z = w
		s.P = s.Z
	}
}

func (s *star) draw(imd *imdraw.IMDraw) {
	p := pixel.V(
		scale(s.X/s.Z, 0, 1, 0, w),
		scale(s.Y/s.Z, 0, 1, 0, h),
	)

	o := pixel.V(
		scale(s.X/s.P, 0, 1, 0, w),
		scale(s.Y/s.P, 0, 1, 0, h),
	)

	r := scale(s.Z, 0, w, 11, 0)

	imd.Color = s.C

	if p.Sub(o).Len() > 6 {
		imd.Push(p, o)
		imd.Line(r)
	}

	imd.Push(p)
	imd.Circle(r, 0)
}

func run() {
	win, err := opengl.NewWindow(opengl.WindowConfig{
		Bounds:      pixel.R(0, 0, w, h),
		VSync:       true,
		Undecorated: true,
	})
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	imd.Precision = 7

	imd.SetMatrix(pixel.IM.Moved(win.Bounds().Center()))

	last := time.Now()

	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixel.KeyEscape) || win.JustPressed(pixel.KeyQ))

		if win.Pressed(pixel.KeyUp) {
			speed += 10
		}

		if win.Pressed(pixel.KeyDown) {
			if speed > 10 {
				speed -= 10
			}
		}

		if win.Pressed(pixel.KeySpace) {
			speed = 100
		}

		d := time.Since(last).Seconds()

		last = time.Now()

		imd.Clear()

		for _, s := range stars {
			s.update(d)
			s.draw(imd)
		}

		win.Clear(color.Black)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	opengl.Run(run)
}

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func scale(unscaledNum, min, max, minAllowed, maxAllowed float64) float64 {
	return (maxAllowed-minAllowed)*(unscaledNum-min)/(max-min) + minAllowed
}

// Colors based on stellar types listed at
// http://www.vendian.org/mncharity/dir3/starcolor/
var Colors = []color.RGBA{
	{157, 180, 255, 255},
	{162, 185, 255, 255},
	{167, 188, 255, 255},
	{170, 191, 255, 255},
	{175, 195, 255, 255},
	{186, 204, 255, 255},
	{192, 209, 255, 255},
	{202, 216, 255, 255},
	{228, 232, 255, 255},
	{237, 238, 255, 255},
	{251, 248, 255, 255},
	{255, 249, 249, 255},
	{255, 245, 236, 255},
	{255, 244, 232, 255},
	{255, 241, 223, 255},
	{255, 235, 209, 255},
	{255, 215, 174, 255},
	{255, 198, 144, 255},
	{255, 190, 127, 255},
	{255, 187, 123, 255},
	{255, 187, 123, 255},
}
