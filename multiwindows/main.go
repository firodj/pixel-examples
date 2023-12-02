package main

import (
	"fmt"

	pixel "github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/text"
)

type EasyWindow1 struct {
	win     *opengl.Window
	txt     *text.Text
	counter int
}

func (w *EasyWindow1) Setup() error {
	w.txt = text.New(pixel.V(0, 0), text.Atlas7x13)

	return nil
}

func (w *EasyWindow1) Win() *opengl.Window {
	return w.win
}

func (w *EasyWindow1) Update() error {
	w.counter++
	return nil
}

func (w *EasyWindow1) Draw() error {
	w.win.Clear(pixel.RGB(0, 0, 0))
	w.txt.Clear()

	fmt.Fprintf(w.txt, "Window 1\n")
	fmt.Fprintf(w.txt, "Counter: %d\n", w.counter)
	fmt.Fprintf(w.txt, "FPS: %.01f\n", wm.FPS())

	w.txt.Draw(w.win, pixel.IM.Scaled(w.txt.Orig, 2))
	return nil
}

type EasyWindow2 struct {
	win     *opengl.Window
	txt     *text.Text
	counter uint64
}

func (w *EasyWindow2) Setup() error {
	w.txt = text.New(pixel.V(0, 0), text.Atlas7x13)
	w.counter = 0
	return nil
}

func (w *EasyWindow2) Win() *opengl.Window {
	return w.win
}

func (w *EasyWindow2) Update() error {
	w.counter--
	return nil
}

func (w *EasyWindow2) Draw() error {
	w.win.Clear(pixel.RGB(0, 0, 0))
	w.txt.Clear()

	fmt.Fprintf(w.txt, "Window 2\n")
	fmt.Fprintf(w.txt, "Counter: %d\n", w.counter)
	fmt.Fprintf(w.txt, "FPS: %.01f\n", wm.FPS())

	w.txt.Draw(w.win, pixel.IM.Scaled(w.txt.Orig, 2))
	return nil
}

var wm *opengl.WindowManager = opengl.NewWindowManager()

func run() {
	w1, err := opengl.NewWindow(opengl.WindowConfig{
		Title:  "Main Window",
		Bounds: pixel.R(0, 0, 200, 200),
	})

	if err != nil {
		panic(err)
	}

	w2, err := opengl.NewWindow(opengl.WindowConfig{
		Title:  "Window 2",
		Bounds: pixel.R(0, 0, 500, 200),
	})

	if err != nil {
		panic(err)
	}

	wm.InsertWindows([]opengl.EasyWindow{
		&EasyWindow1{win: w1},
		&EasyWindow2{win: w2},
	})

	wm.SetFPS(60)

	if err := wm.Loop(); err != nil {
		panic(err)
	}

}

func main() {
	opengl.Run(run)
}
