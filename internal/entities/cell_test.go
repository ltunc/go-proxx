package entities

import "testing"

func TestCell_NearHoles(t *testing.T) {
	type fields struct {
		IsOpen      bool
		IsBlackHole bool
		neighbors   []*Cell
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"none",
			fields{
				IsOpen:      false,
				IsBlackHole: false,
				neighbors: []*Cell{
					{IsOpen: false, IsBlackHole: false, neighbors: nil},
					{IsOpen: true, IsBlackHole: false, neighbors: nil},
				},
			},
			0,
		},
		{
			"some",
			fields{
				IsOpen:      false,
				IsBlackHole: false,
				neighbors: []*Cell{
					{IsOpen: false, IsBlackHole: false, neighbors: nil},
					{IsOpen: true, IsBlackHole: true, neighbors: nil},
					{IsOpen: true, IsBlackHole: false, neighbors: nil},
					{IsOpen: true, IsBlackHole: true, neighbors: nil},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				IsOpen:      tt.fields.IsOpen,
				IsBlackHole: tt.fields.IsBlackHole,
				neighbors:   tt.fields.neighbors,
			}
			if got := c.NearHoles(); got != tt.want {
				t.Errorf("NearHoles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_click(t *testing.T) {
	type fields struct {
		IsOpen      bool
		IsBlackHole bool
		neighbors   []*Cell
	}
	tests := []struct {
		name              string
		fields            fields
		want              bool
		wantNeighborsOpen bool
	}{
		{
			"safe",
			fields{
				IsOpen:      false,
				IsBlackHole: false,
				neighbors: []*Cell{
					{IsOpen: false, IsBlackHole: false, neighbors: nil},
					{IsOpen: false, IsBlackHole: true, neighbors: nil},
					{IsOpen: false, IsBlackHole: false, neighbors: nil},
					{IsOpen: false, IsBlackHole: true, neighbors: nil},
				},
			},
			true,
			false,
		},
		{
			"black hole",
			fields{
				IsOpen:      false,
				IsBlackHole: true,
				neighbors: []*Cell{
					{IsOpen: false, IsBlackHole: false, neighbors: nil},
					{IsOpen: false, IsBlackHole: true, neighbors: nil},
					{IsOpen: false, IsBlackHole: false, neighbors: nil},
					{IsOpen: false, IsBlackHole: true, neighbors: nil},
				},
			},
			false,
			false,
		},
		{
			"no blackholes",
			fields{
				IsOpen:      false,
				IsBlackHole: false,
				neighbors: []*Cell{
					{IsOpen: false, IsBlackHole: false, neighbors: nil},
					{IsOpen: false, IsBlackHole: false, neighbors: nil},
				},
			},
			true,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				IsOpen:      tt.fields.IsOpen,
				IsBlackHole: tt.fields.IsBlackHole,
				neighbors:   tt.fields.neighbors,
			}
			if got := c.click(); got != tt.want {
				t.Errorf("click() = %v, want %v", got, tt.want)
			}
			if !c.IsBlackHole && !c.IsOpen {
				t.Errorf("click() IsOpen is false, want true")
			}
			if tt.wantNeighborsOpen {
				for _, n := range c.neighbors {
					if !n.IsOpen {
						t.Errorf("click() didn't opened neighbors")
					}
				}
			}
		})
	}
}
