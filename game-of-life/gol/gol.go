package gol

import (
	"fmt"
	tm "github.com/buger/goterm"
	"math/rand"
	"time"
)

type state int

const (
	KILLED = iota
	ACTIVE
)

func setupWorld(size int) [][]state {
	world := [][]state{}
	for i := 0; i < size; i++ {
		row := []state{}
		for j := 0; j < size; j++ {
			row = append(row, resolveState(rand.Intn(2)))
		}
		world = append(world, row)
	}
	return world
}
func resolveState(value int) state {
	if value == 0 {
		return KILLED
	}
	return ACTIVE
}

//Gol main loop
func Gol(size int, generations int) {
	world := setupWorld(size)
	tm.Clear()
	for i := 0; i < generations; i++ {
		tm.MoveCursor(1,1)
		world = recalculate(world)
		fmt.Println(i)
		printWorld(world)
		tm.Flush()
		time.Sleep(time.Second)
	}
}

func recalculate(world [][]state) [][]state {
	var size int = len(world)
	newWorld := make([][]state, size)
	for x := range newWorld {
		newWorld[x] = make([]state, size)
		for y := range newWorld[x] {
			var currentState = world[x][y]
			var newState state = currentState
			var liveNeighbours = countLiveNeighbours(x, y, world)
			if currentState == ACTIVE {
				//A cell with fewer than two live neighbours dies of under-population
				if liveNeighbours < 2 {
					newState = KILLED
				}
				//A cell with 2 or 3 live neighbours lives on to the next generation
				if liveNeighbours == 2 || liveNeighbours == 3 {
					newState = ACTIVE
				}
				if liveNeighbours > 3 {
					newState = KILLED
				}
			}
			if currentState == KILLED && liveNeighbours == 3 {
				newState = ACTIVE
			}
			newWorld[x][y] = newState
		}
	}
	return newWorld
}

func printWorld(world [][]state) {
	for _, x := range world {
		fmt.Println(x)
	}
}
func countLiveNeighbours(x int, y int, world [][]state) int {
	var size int = len(world)
	liveNeighbours := 0
	left := x - 1
	if left < 0 {
		left = size - 1
	}
	if world[left][y] == ACTIVE {
		liveNeighbours++
	}

	right := x + 1
	if right == size {
		right = 0
	}
	if world[right][y] == ACTIVE {
		liveNeighbours++
	}

	above := y - 1
	if above < 0 {
		above = size - 1
	}
	if world[x][above] == ACTIVE {
		liveNeighbours++
	}
	below := y + 1
	if below == size {
		below = 0
	}
	if world[x][below] == ACTIVE {
		liveNeighbours++
	}
	return liveNeighbours
}
