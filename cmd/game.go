package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	// create empty board with requested size
	b := createBoard(4, 4)
	// link all neighbors, to allow easy calculation and operations with neighbors
	linkNeighbors(b)
	// add mines to the board
	populate(b, 3)
	// render the board to the console
	render(b)
	reader := bufio.NewReader(os.Stdin)
	for {
		t, _ := reader.ReadString('\n')
		t = strings.Replace(t, "\n", "", -1)
		coordinates := strings.Split(t, ",")
		if len(coordinates) != 2 {
			fmt.Println("wrong format, expected coordinates x,y")
			continue
		}
		x, err := strconv.Atoi(coordinates[0])
		if err != nil {
			fmt.Println("the x is not an integer")
			continue
		}
		y, err := strconv.Atoi(coordinates[1])
		if err != nil {
			fmt.Println("the y is not an integer")
			continue
		}
		if err := b.click(x, y); err != nil {
			fmt.Printf("error: %v", err)
			continue
		}
		fmt.Println("\n====")
		render(b)
	}
}

func (b board) click(x, y int) error {
	c, ok := getCell(b, x, y)
	if !ok {
		return fmt.Errorf("cannot click on the cell, wrong coordinates")
	}
	if !c.click() {
		fmt.Printf("you loose.")
	}
	return nil
}

func (c *cell) click() bool {
	if c.isOpen {
		return true
	}
	if c.isBlackHole {
		return false
	}
	c.isOpen = true
	if c.nearHoles() == 0 {
		for _, n := range c.neighbors {
			n.click()
		}
	}
	return true
}

func render(b board) {
	for i := -1; i < len(b[0]); i++ {
		if i < 0 {
			fmt.Printf("  |")
			continue
		}
		fmt.Printf("%2d.|", i)
	}
	fmt.Printf("\n%s\n", strings.Repeat("-", len(b[0])*4+3))
	for i, row := range b {
		fmt.Printf("%d.|", i)
		for _, c := range row {
			v := " "
			if c.isOpen {
				if c.isBlackHole {
					v = "H"
				} else {
					v = strconv.Itoa(c.nearHoles())
				}
			}
			fmt.Printf("%2s |", v)
		}

		fmt.Printf("\n%s\n", strings.Repeat("-", len(b[0])*4+3))
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
