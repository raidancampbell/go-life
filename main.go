package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)
const (
	GRID_WIDTH = 20
	GRID_HEIGHT = 20
	WINDOW_WIDTH = 1024
	WINDOW_HEIGHT = 1024
)
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "go-life",
		Bounds: pixel.R(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	w := newWorld(GRID_WIDTH,GRID_HEIGHT)
	w.init()

	frames:=0
	second := time.Tick(time.Second)
	for !win.Closed() {
		// safe to remove: vsync is set.  60fps is too fast for viewing though.
		time.Sleep(35 * time.Millisecond)
		win.Clear(colornames.Black)
		imd.Clear()

		drawGrid(cfg, imd)
		drawCells(w, imd, cfg)

		if win.Pressed(pixelgl.KeySpace) || win.Pressed(pixelgl.KeyEnter) ||win.Pressed(pixelgl.MouseButtonLeft) {
			w.init()
		}

		imd.Draw(win)
		win.Update()
		w.tick()
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

func drawGrid(cfg pixelgl.WindowConfig, imd *imdraw.IMDraw) {
	// horizontal grid
	for i := 0; i <= GRID_WIDTH; i++ {
		// calculate grid width
		width := (cfg.Bounds.Max.Y - cfg.Bounds.Min.Y)/GRID_WIDTH
		imd.Push(pixel.V(cfg.Bounds.Min.X, width * float64(i)), pixel.V(cfg.Bounds.Max.X, width * float64(i)))
		imd.Line(5)
	}

	// vertical grid
	for i := 0; i <= GRID_HEIGHT; i++ {
		// calculate grid height
		height := (cfg.Bounds.Max.X - cfg.Bounds.Min.X)/GRID_HEIGHT
		imd.Push(pixel.V(height * float64(i), cfg.Bounds.Min.Y), pixel.V(height * float64(i), cfg.Bounds.Max.Y))
		imd.Line(5)
	}
}

func drawCells(w world, imd *imdraw.IMDraw, cfg pixelgl.WindowConfig) {
	for x := 0; x <= GRID_WIDTH; x++ {
		for y := 0; y <= GRID_HEIGHT; y++ {
			if w.isAlive(x, y) {
				centerx := ((cfg.Bounds.Max.X - cfg.Bounds.Min.X)/GRID_WIDTH) * float64(x) + ((cfg.Bounds.Max.X - cfg.Bounds.Min.X)/GRID_WIDTH)/2
				centery := ((cfg.Bounds.Max.Y - cfg.Bounds.Min.Y)/GRID_HEIGHT) * float64(y) + ((cfg.Bounds.Max.X - cfg.Bounds.Min.X)/GRID_HEIGHT)/2
				imd.Push(pixel.V(centerx, centery))
				imd.Circle(((cfg.Bounds.Max.Y - cfg.Bounds.Min.Y)/GRID_HEIGHT)/3, 0)
			}
		}
	}
}

func main() {
	pixelgl.Run(run)
}