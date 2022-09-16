# Proxx game

The programm consist of 2 main entities:

* Cell -- represent a cell, may contain a black hole
* Board -- represent the game board, consists of cells. Is a 2D slice of pointers to Cell objects

After creating a board method Board.LinkNeighbors() must be called to prepare cells.
This method finds and stores all neighbors of the cell.
It simplifies manipulation with cells when user clicks on one of them,
allows to easily calculate nearest blackholes,
open neighbors if there are no blackholes nearby.

Method Board.Populate() places blackholes on the board (set flag IsBlackhole for some cells).
The distribution of the holes **is not normal** (was not enough time to proper implement normal distribution), instead just random.
To place exact amount of holes the method uses linear representation (1d) of 2d array, sets for K items flag true (this represents a blackhole),
shuffles this linear representation, then converts items in 2d coordinates.
Another approach is to generate X and Y coordinates randomly, but chosen method of shuffle allows to potentially easily implementation of normal distribution.

Methods in cmd/game.go main.userInput() and main.render() only for demonstration, they not handle all cases, but simple shows now it may work all together

## How to play:

    go run cmd/game.go

then enter coordinates: X,Y
they must be integers, separated by coma, without spaces


