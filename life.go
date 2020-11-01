package main

import (
	"math/rand"
	"time"
)

type world struct {
	w[][] bool
}

func newWorld(width, height int) world {
	w := make([][]bool, width)
	for i := range w {
		w[i] = make([]bool, height)
	}
	return world{w: w}
}

func (w world) init() {
	rand.Seed(time.Now().UnixNano())
	for rownum, row := range w.w {
		for colnum := range row{
			w.w[rownum][colnum] = rand.Int() % 2 == 0
		}
	}
}

/*
Any live cell with fewer than two live neighbours dies, as if by underpopulation.
Any live cell with two or three live neighbours lives on to the next generation.
Any live cell with more than three live neighbours dies, as if by overpopulation.
Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
*/
func (w world) tick() {
	for rownum, row := range w.w {
		for colnum := range row{
			neighbors := w.getNeighbors(rownum, colnum)
			// killed by underpopulation
			if neighbors < 2 {
				w.w[rownum][colnum] = false
			}

			// killed by overpopulation
			if neighbors > 3 {
				w.w[rownum][colnum] = false
			}

			// birth
			if neighbors == 3 {
				w.w[rownum][colnum] = true
			}
		}
	}
}

func (w world) getNeighbors(x, y int) int {
	return sumTrue(w.isAlive(x+1,y),
		w.isAlive(x+1,y+1),
		w.isAlive(x,y+1),
		w.isAlive(x-1,y+1),
		w.isAlive(x-1,y),
		w.isAlive(x-1,y-1),
		w.isAlive(x,y-1),
		w.isAlive(x+1,y-1))
}

// isAlive provides a bound-safe answer to whether the given cell is alive
// if the given cell is outside the grid, it is assumed to be dead
func (w world) isAlive(x,y int) bool {
	if x < 0 || x >= len(w.w) {
		return false
	}
	if y < 0 || y >= len(w.w[0]) {
		return false
	}
	return w.w[x][y]
}

func sumTrue(b ...bool) int {
	sum := 0
	for _, val := range b {
		if val {
			sum += 1
		}
	}
	return sum
}