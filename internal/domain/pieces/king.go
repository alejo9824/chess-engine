package pieces

import "chess/internal/domain"

type King struct {
	X     int
	Y     int
	Color domain.Color `json:"color"`
	Name  string       `json:"name"`
}

func NewKing(color domain.Color) *King {
	return &King{
		Color: color,
		Name:  "King",
	}
}

func (p *King) Move(x int, y int, x2 int, y2 int, targetEmpty bool) bool {

	if (x2 == x+1 || x2 == x-1 || x2 == x) && (y2 == y+1 || y2 == y-1 || y2 == y) {
		return true
	}
	return false

}

func (p *King) GetColor() domain.Color {
	return p.Color
}
