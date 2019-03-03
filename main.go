package main

import (
	"fmt"
	"math/rand"
	"time"
)

// mazes
const (
	H              = 40
	W              = 80
	wallVertical   = "|"
	wallHorizontal = "-"
)

var mazeMap = newMap()

func newMap() [][]string {
	mm := make([][]string, H)
	for i := 0; i < H; i++ {
		mm[i] = make([]string, W)
	}
	return mm
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			mazeMap[y][x] = " "
		}
	}

	addRoom(true)
	for i := 0; i < 5000; i++ {
		addRoom(false)
	}

	for y := range mazeMap {
		for x := range mazeMap[y] {
			fmt.Print(mazeMap[y][x])
		}
		fmt.Print("\n")
	}
}

func addRoom(start bool) {
	w := rand.Intn(10) + 5
	h := rand.Intn(6) + 3
	rx := rand.Intn(W-w-2) + 1
	ry := rand.Intn(H-h-2) + 1

	// see if it's blocked or allowed
	for y := ry - 1; y < ry+h+2; y++ {
		for x := rx - 1; x < rx+w+2; x++ {
			if mazeMap[y][x] == "." {
				return
			}
		}
	}

	doorCount := 0
	var dx, dy int

	if !start {
		// canPlace := false
		for x := rx; x < rx+w; x++ {
			if mazeMap[ry-1][x] == wallVertical || mazeMap[ry-1][x] == wallHorizontal {
				// canPlace = true
				doorCount++
				if rand.Intn(doorCount) == 0 {
					dx = x
					dy = ry - 1
				}
			}
			if mazeMap[ry+h][x] == wallVertical || mazeMap[ry+h][x] == wallHorizontal {
				// canPlace = true
				doorCount++
				if rand.Intn(doorCount) == 0 {
					dx = x
					dy = ry + h
				}
			}
		}

		for y := ry; y < ry+h; y++ {
			if mazeMap[y][rx-1] == wallVertical || mazeMap[y][rx-1] == wallHorizontal {
				// canPlace = true
				doorCount++
				if rand.Intn(doorCount) == 0 {
					dx = rx - 1
					dy = y
				}
			}
			if mazeMap[y][rx+w] == wallVertical || mazeMap[y][rx+w] == wallHorizontal {
				// canPlace = true
				doorCount++
				if rand.Intn(doorCount) == 0 {
					dx = rx + w
					dy = y
				}
			}
		}

		if doorCount == 0 {
			return
		}
	}

	for y := ry; y < ry+h; y++ {
		for x := rx; x < rx+w; x++ {
			mazeMap[y][x] = "."
		}
	}

	for x := rx; x < rx+w; x++ {
		mazeMap[ry-1][x] = wallHorizontal
		mazeMap[ry+h][x] = wallHorizontal
	}

	for y := ry; y < ry+h; y++ {
		mazeMap[y][rx-1] = wallVertical
		mazeMap[y][rx+w] = wallVertical
	}

	if doorCount > 0 {
		mazeMap[dy][dx] = "+"
	}

	var startX = rand.Intn(6) + 1
	if start {
		startX = 1
	}
	for i := 0; i < startX; i++ {
		thing := string(rune(65 + rand.Intn(62)))
		if start {
			thing = "@"
		} else if rand.Intn(4) == 0 {
			thing = "$"
		}
		mazeMap[rand.Intn(h)+ry][rand.Intn(w)+rx] = thing
	}
}
