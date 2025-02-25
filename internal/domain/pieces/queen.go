package pieces

import "chess/internal/domain"

type Queen struct {
	X     int
	Y     int
	Color domain.Color `json:"color"`
	Name  string       `json:"name"`
}

func NewQueen(color domain.Color) *Queen {
	return &Queen{
		Color: color,
		Name:  "Queen",
	}
}

func (p *Queen) Move(x int, y int, x2 int, y2 int, targetEmpty bool) bool {

	if (x2 >= x && y == y2) || (y2 >= y && x == x2) {

		return true
	}
	if (x2 < x && y == y2) || (y2 < y && x == x2) {

		return true
	}

	if abs(x2-x) == abs(y2-y) {
		return true
	}

	return false
}

func (p *Queen) GetColor() domain.Color {
	return p.Color
}

func (p *Queen) GetName() string {
	return p.Name
}
