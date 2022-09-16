package entities

import (
	"reflect"
	"testing"
)

func TestBoard_Click(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		b       Board
		args    args
		want    bool
		wantErr bool
	}{
		{
			"safe",
			Board{{&Cell{IsOpen: false, IsBlackHole: false, neighbors: nil}}},
			args{0, 0},
			true,
			false,
		},
		{
			"black hole",
			Board{{&Cell{IsOpen: false, IsBlackHole: true, neighbors: nil}}},
			args{0, 0},
			false,
			false,
		},
		{
			"out of range",
			Board{{&Cell{IsOpen: false, IsBlackHole: false, neighbors: nil}}},
			args{9, 0},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.Click(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("Click() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Click() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_GetCell(t *testing.T) {
	type args struct {
		x int
		y int
	}
	demoBoard := Board{
		{&Cell{IsOpen: false, IsBlackHole: false, neighbors: nil}, &Cell{IsOpen: false, IsBlackHole: true, neighbors: nil}},
		{&Cell{IsOpen: false, IsBlackHole: false, neighbors: nil}},
	}
	tests := []struct {
		name  string
		b     Board
		args  args
		want  *Cell
		want1 bool
	}{
		{
			"simple",
			demoBoard,
			args{0, 1},
			&Cell{IsOpen: false, IsBlackHole: false, neighbors: nil},
			true,
		},
		{
			"y out of range",
			demoBoard,
			args{0, 2},
			nil,
			false,
		},
		{
			"x out of range",
			demoBoard,
			args{2, 0},
			nil,
			false,
		},
		{
			"negative coordinates",
			demoBoard,
			args{2, -1},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.b.getCell(tt.args.x, tt.args.y)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCell() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getCell() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBoard_LinkNeighbors(t *testing.T) {
	cells := []*Cell{
		{IsOpen: false, IsBlackHole: false, neighbors: nil},
		{IsOpen: true, IsBlackHole: true, neighbors: nil},
		{IsOpen: false, IsBlackHole: false, neighbors: nil},
		{IsOpen: false, IsBlackHole: true, neighbors: nil},
		{IsOpen: false, IsBlackHole: false, neighbors: nil},
		{IsOpen: true, IsBlackHole: true, neighbors: nil},
		{IsOpen: false, IsBlackHole: true, neighbors: nil},
		{IsOpen: true, IsBlackHole: false, neighbors: nil},
		{IsOpen: false, IsBlackHole: true, neighbors: nil},
	}
	demoBoard := Board{
		{cells[0], cells[1], cells[2]},
		{cells[3], cells[4], cells[5]},
		{cells[6], cells[7], cells[8]},
	}
	tests := []struct {
		name  string
		b     Board
		wantX int
		wantY int
		want  []*Cell
	}{
		{
			"board 3x3 cell 0x0",
			demoBoard,
			0,
			0,
			[]*Cell{cells[1], cells[3], cells[4]},
		},
		{
			"board 3x3 cell 1x1",
			demoBoard,
			1,
			1,
			[]*Cell{cells[0], cells[1], cells[2], cells[3], cells[5], cells[6], cells[7], cells[8]},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.LinkNeighbors()
			got := tt.b[tt.wantY][tt.wantX].neighbors
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkNeighbors() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_OpenAll(t *testing.T) {
	cells := []*Cell{
		{IsOpen: false, IsBlackHole: false, neighbors: nil},
		{IsOpen: true, IsBlackHole: true, neighbors: nil},
		{IsOpen: false, IsBlackHole: false, neighbors: nil},
		{IsOpen: false, IsBlackHole: true, neighbors: nil},
	}
	demoBoard := Board{
		{cells[0], cells[1]},
		{cells[2], cells[3]},
	}
	tests := []struct {
		name string
		b    Board
	}{
		{
			"basic",
			demoBoard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.OpenAll()
			for y, row := range tt.b {
				for x, c := range row {
					if !c.IsOpen {
						t.Errorf("OpenAll() cell %dx%d is closed", x, y)
					}
				}
			}
		})
	}
}

func TestBoard_Populate(t *testing.T) {
	type args struct {
		k int
	}
	cells := []*Cell{
		{IsOpen: false, IsBlackHole: false, neighbors: nil},
		{IsOpen: true, IsBlackHole: false, neighbors: nil},
		{IsOpen: false, IsBlackHole: false, neighbors: nil},
		{IsOpen: false, IsBlackHole: false, neighbors: nil},
	}
	demoBoard := Board{
		{cells[0], cells[1]},
		{cells[2], cells[3]},
	}
	tests := []struct {
		name    string
		b       Board
		args    args
		wantErr bool
	}{
		{
			"board 2x2, 2 holes",
			demoBoard,
			args{2},
			false,
		},
		{
			"board 2x2, 9 holes",
			demoBoard,
			args{9},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// reset all cells
			for _, c := range cells {
				c.IsBlackHole = false
			}
			err := tt.b.Populate(tt.args.k)
			if tt.wantErr {
				if (err != nil) != tt.wantErr {
					t.Errorf("Populate() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				cnt := 0
				for _, c := range cells {
					if c.IsBlackHole {
						cnt++
					}
				}
				if tt.args.k != cnt {
					t.Errorf("Populate() unexpected number of blackholes: got %v, want %v", cnt, tt.args.k)
				}
			}
		})
	}
}

func TestNewBoard(t *testing.T) {
	type args struct {
		w int
		h int
	}
	tests := []struct {
		name string
		args args
		//want Board
		wantYLen int
		wantXLen int
	}{
		{
			"2x2",
			args{w: 2, h: 2},
			2,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBoard(tt.args.w, tt.args.h)
			if yLen := len(got); yLen != tt.wantYLen {
				t.Errorf("NewBoard() got Y length %v, want %v", yLen, tt.wantYLen)
			}
			for y, row := range got {
				if xLen := len(row); xLen != tt.wantXLen {
					t.Errorf("NewBoard(), row %v, got X length %v, want %v", y, xLen, tt.wantXLen)
				}

			}
		})
	}
}
