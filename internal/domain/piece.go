package domain

type Piece interface {
	GetColor() Color
	Move(x int, y int, x2 int, y2 int, targetEmpty bool) bool
}

type Color string

const (
	White Color = "white"
	Black Color = "black"
)
