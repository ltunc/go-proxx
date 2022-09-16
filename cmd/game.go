package main

import (
	"bufio"
	"fmt"
	"os"
	"proxx/internal/entities"
	"strconv"
	"strings"
)

func main() {
	// create empty board with requested size
	b := entities.NewBoard(4, 4)
	// link all neighbors, to allow easy calculation and operations with neighbors
	b.LinkNeighbors()
	// add mines to the board
	if err := b.Populate(3); err != nil {
		panic(err)
	}
	// render the board to the console
	render(b)
	userInput(b, bufio.NewReader(os.Stdin))
}

func userInput(b entities.Board, reader *bufio.Reader) {
	for {
		str, _ := reader.ReadString('\n')
		str = strings.Replace(str, "\n", "", -1)
		coordinates := strings.Split(str, ",")
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
		isSafe, err := b.Click(x, y)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		if !isSafe {
			b.OpenAll()
			render(b)
			fmt.Println("\nYou lose!")
			break
		}
		fmt.Println("\n====")
		render(b)
	}
}

// render prints board to console
// for debug and visual representation
func render(b entities.Board) {
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
			if c.IsOpen {
				if c.IsBlackHole {
					v = "H"
				} else {
					v = strconv.Itoa(c.NearHoles())
				}
			}
			fmt.Printf("%2s |", v)
		}
		fmt.Printf("\n%s\n", strings.Repeat("-", len(b[0])*4+3))
	}
}
