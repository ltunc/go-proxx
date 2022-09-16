package entities

import (
	"fmt"
	"math/rand"
	"time"
)

type Board [][]*Cell

// x x x
// y
// y

// NewBoard creates a board with requested dimensions
func NewBoard(w, h int) Board {
	b := make(Board, w)
	for i := 0; i < h; i++ {
		b[i] = make([]*Cell, w)
	}
	for y, row := range b {
		for x := range row {
			b[y][x] = &Cell{IsOpen: false, IsBlackHole: false, neighbors: []*Cell{}}
		}
	}
	return b
}

// GetCell finds and returns a cell by its coordinates on the board
// returns false if cell not exist (coordinates are outside the board)
func (b Board) GetCell(x, y int) (*Cell, bool) {
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

// LinkNeighbors creates links between neighboring cells on the board
// adds neighbors to cell.neighbors slice of each cell
func (b Board) LinkNeighbors() {
	for y, row := range b {
		for x, c := range row {
			// up left
			if p, ok := b.GetCell(x-1, y-1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// up
			if p, ok := b.GetCell(x, y-1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// up right
			if p, ok := b.GetCell(x+1, y-1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// left
			if p, ok := b.GetCell(x-1, y); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// right
			if p, ok := b.GetCell(x+1, y); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// down left
			if p, ok := b.GetCell(x-1, y+1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// down
			if p, ok := b.GetCell(x, y+1); ok {
				c.neighbors = append(c.neighbors, p)
			}
			// down right
			if p, ok := b.GetCell(x+1, y+1); ok {
				c.neighbors = append(c.neighbors, p)
			}
		}
	}
}

// Populate adds k black holes to the board
func (b Board) Populate(k int) error {
	xLen := len(b[0])
	yLen := len(b)
	linear := make([]bool, xLen*yLen)
	if k > len(linear) {
		return fmt.Errorf("number of mines is greater than amount of cells")
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
		b[originalY][originalX].IsBlackHole = isHole
	}
	return nil
}

// Click processes user's input (click)
// returns true if click was successful and not on a black hole,
// returns false if the user clicked on a black hole
func (b Board) Click(x, y int) (bool, error) {
	c, ok := b.GetCell(x, y)
	if !ok {
		return false, fmt.Errorf("cannot click on the cell, wrong coordinates")
	}
	if !c.click() {
		return false, nil
	}
	return true, nil
}

// OpenAll marks all cells of the board as open
func (b Board) OpenAll() {
	for _, row := range b {
		for _, c := range row {
			c.IsOpen = true
		}
	}
}
