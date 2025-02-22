package pieces

import "chess/internal/domain"

type Rook struct {
	X     int
	Y     int
	Color domain.Color
	Name  string `json:"name"`
}

func NewRook(color domain.Color) *Rook {
	return &Rook{
		Color: color,
		Name:  "Rook",
	}
}

func (p *Rook) Move(x int, y int, x2 int, y2 int, targetEmpty bool) bool {

	if (x2 >= x && y == y2) || (y2 >= y && x == x2) {

		return true
	}
	if (x2 < x && y == y2) || (y2 < y && x == x2) {

		return true
	}

	return false

}

func (p *Rook) GetColor() domain.Color {
	return p.Color
}
