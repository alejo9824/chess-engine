package domain

type Piece interface {
	GetName() string
	GetColor() Color
	Move(x int, y int, x2 int, y2 int, targetEmpty bool) bool
}

type Color string
type Name string

const (
	White  Color = "white"
	Black  Color = "black"
	Queen  Name  = "Queen"
	King   Name  = "King"
	Rook   Name  = "Rook"
	Bishop Name  = "Bishop"
	Knight Name  = "Knight"
	Pawn   Name  = "Pawn"
)
