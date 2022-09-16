package entities

type Cell struct {
	IsOpen      bool
	IsBlackHole bool
	neighbors   []*Cell
}

// NearHoles counts amount of black holes neighboring the cell
func (c *Cell) NearHoles() int {
	cnt := 0
	for _, e := range c.neighbors {
		if e.IsBlackHole {
			cnt++
		}
	}
	return cnt
}

// click processes logic of clicking on the cell
// returns true if click is successful and the cell does not contain a black hole
func (c *Cell) click() bool {
	if c.IsOpen {
		return true
	}
	if c.IsBlackHole {
		return false
	}
	c.IsOpen = true
	if c.NearHoles() == 0 {
		for _, n := range c.neighbors {
			n.click()
		}
	}
	return true
}
