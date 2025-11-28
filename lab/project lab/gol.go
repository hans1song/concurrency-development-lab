// Package main implements a parallel version of Conway's Game of Life using the Ebiten library.
// It demonstrates data parallelism by dividing the grid update task among multiple worker goroutines.
package main

import (
	"image/color"
	"log"
	"math/rand"
	"sync"

	"github.com/hajimehoshi/ebiten"
)

// scale determines the pixel size of each cell.
const scale int = 2

// width and height define the dimensions of the grid.
const width = 300
const height = 300

// numWorkers is the number of goroutines running in parallel to update the grid.
const numWorkers = 8

var blue color.Color = color.RGBA{69, 145, 196, 255}
var yellow color.Color = color.RGBA{255, 230, 120, 255}

// grid holds the current state of the game (1 for alive, 0 for dead).
var grid [width][height]uint8 = [width][height]uint8{}

// buffer is used to calculate the next state without mutating the current grid during calculation.
var buffer [width][height]uint8 = [width][height]uint8{}

var count int = 0

// update calculates the next generation of the Game of Life.
// It splits the grid into horizontal strips and assigns each strip to a worker goroutine.
func update() error {
	var wg sync.WaitGroup

	rowsPerWorker := (height - 2) / numWorkers

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		// Calculate the range of rows this worker handles.
		startY := 1 + w*rowsPerWorker
		endY := startY + rowsPerWorker

		// Ensure the last worker handles any remaining rows.
		if w == numWorkers-1 {
			endY = height - 1
		}

		go func(startRow, endRow int) {
			defer wg.Done()

			for x := 1; x < width-1; x++ {
				for y := startRow; y < endRow; y++ {
					// Count live neighbors
					n := grid[x-1][y-1] + grid[x-1][y+0] + grid[x-1][y+1] +
						grid[x+0][y-1] + grid[x+0][y+1] +
						grid[x+1][y-1] + grid[x+1][y+0] + grid[x+1][y+1]

					// Apply Game of Life rules
					if grid[x][y] == 0 && n == 3 {
						buffer[x][y] = 1 // Reproduction
					} else if n < 2 || n > 3 {
						buffer[x][y] = 0 // Underpopulation or Overpopulation
					} else {
						buffer[x][y] = grid[x][y] // Survival
					}
				}
			}
		}(startY, endY)
	}

	// Wait for all workers to finish updating their sections.
	wg.Wait()

	// Swap the buffers: buffer becomes the current grid for the next frame.
	temp := buffer
	buffer = grid
	grid = temp
	return nil
}

func display(window *ebiten.Image) {
	window.Fill(blue)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for i := 0; i < scale; i++ {
				for j := 0; j < scale; j++ {
					if grid[x][y] == 1 {
						window.Set(x*scale+i, y*scale+j, yellow)
					}
				}
			}
		}
	}
}

func frame(window *ebiten.Image) error {
	count++
	var err error = nil
	if count == 20 {
		err = update()
		count = 0
	}
	if !ebiten.IsDrawingSkipped() {
		display(window)
	}

	return err
}

func main() {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if rand.Float32() < 0.5 {
				grid[x][y] = 1
			}
		}
	}

	if err := ebiten.Run(frame, width, height, 2, "Game of Life"); err != nil {
		log.Fatal(err)
	}
}
