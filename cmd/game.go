package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type cell struct {
	isOpen      bool
	isBlackHole bool
	neighbors   []*cell
}

func (c *cell) nearHoles() int {
	cnt := 0
	for _, e := range c.neighbors {
		if e.isBlackHole {
			cnt++
		}
	}
	return cnt
}

type board [][]*cell

// x x x
// y
// y

func main() {
	b := createBoard(8, 8)
	linkNeighbors(b)
	populate(b, 12)
	render(b)
}

func render(b board) {
	for _, row := range b {
		for _, c := range row {
			v := ""
			if c.isBlackHole {
				v = "H"
			} else {
				if cnt := c.nearHoles(); cnt > 0 {
					v = strconv.Itoa(cnt)
				}
			}
			fmt.Printf("\t%s", v)
		}
		fmt.Print("\n")
	}
}

func createBoard(w, h int) board {
	b := make(board, w)
	for i := 0; i < h; i++ {
		b[i] = make([]*cell, w)
	}
	for y, row := range b {
		for x := range row {
			b[y][x] = &cell{isOpen: false, isBlackHole: false, neighbors: []*cell{}}
		}
	}
	return b
}

// linkNeighbors creates links between neighboring cells on the board
// adds neighbors to cell.neighbors slice of each cell
func linkNeighbors(b board) {
	for y, row := range b {
		for x, c := range row {
			// up left
			if p, ok := getCell(b, x-1, y-1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// up
			if p, ok := getCell(b, x, y-1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// up right
			if p, ok := getCell(b, x+1, y-1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// left
			if p, ok := getCell(b, x-1, y); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// right
			if p, ok := getCell(b, x+1, y); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// down left
			if p, ok := getCell(b, x-1, y+1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// down
			if p, ok := getCell(b, x, y+1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// down right
			if p, ok := getCell(b, x+1, y+1); ok {
				c.neighbors = append(c.neighbors, p)
			}
		}
	}
}

// getCell finds and returns a cell by its coordinates on the board
// returns false if cell not exist (coordinates are outside the board)
func getCell(b board, x, y int) (*cell, bool) {
	if x < 0 || y < 0 {
		return nil, false
	}
	if y >= len(b) {
		return nil, false
	}
	if x >= len(b[y]) {
		return nil, false
	}
	return b[y][x], true
}

// populate adds k black holes to the board
func populate(b board, k int) {
	xLen := len(b[0])
	yLen := len(b)
	linear := make([]bool, xLen*yLen)
	if k > len(linear) {
		panic("number of mines is greater than amount of cells")
	}
	for i := 0; i < k; i++ {
		linear[i] = true
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixMilli()))
	// TODO: use NormFloat64()
	rnd.Shuffle(len(linear), func(i, j int) {
		linear[i], linear[j] = linear[j], linear[i]
	})
	// apply 1D representation to 2D
	for i, isHole := range linear {
		originalX := i % xLen
		originalY := i / yLen
		b[originalY][originalX].isBlackHole = isHole
	}
}
