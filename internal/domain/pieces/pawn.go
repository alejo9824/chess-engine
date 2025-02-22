package pieces

import (
	"chess/internal/domain"
	"log"
)

type Pawn struct {
	X     int
	Y     int
	Color domain.Color
	Name  string `json:"name"`
}

func NewPawn(color domain.Color) *Pawn {
	return &Pawn{
		Color: color,
		Name:  "pawn",
	}
}

func (p *Pawn) Move(x int, y int, x2 int, y2 int, targetEmpty bool) bool {

	if p.Color == domain.White {
		log.Print("primer if color")
		if x == 1 {
			log.Print("primer if")
			if (x2 == x+1 || x2 == 3) && y == y2 {
				log.Print("segundo if")
				return true
			}
		}
		if x2 == x+1 && y == y2 {
			return true
		}
		if (x2 == x+1 && (y2 == y+1 || y2 == y-1)) && !targetEmpty {
			return true
		}

	} else {
		if x == 6 {
			if (x2 == x-1 || x2 == 4) && y == y2 {
				return true
			}
		}
		if x2 == x-1 && y == y2 {
			return true
		}
		if (x2 == x-1 && (y2 == y+1 || y2 == y-1)) && !targetEmpty {
			return true
		}

	}
	return false
}

func (p *Pawn) GetColor() domain.Color {
	return p.Color
}
