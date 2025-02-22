package pieces

import "chess/internal/domain"

type Bishop struct {
	X     int
	Y     int
	Color domain.Color `json:"color"`
	Name  string       `json:"name"`
}

func NewBishop(color domain.Color) *Bishop {
	return &Bishop{
		Color: color,
		Name:  "Bishop",
	}
}

func (p *Bishop) Move(x int, y int, x2 int, y2 int, targetEmpty bool) bool {

	if abs(x2-x) == abs(y2-y) {
		return true
	}
	return false

}

func (p *Bishop) GetColor() domain.Color {
	return p.Color
}

// Función auxiliar para calcular el valor absoluto
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
