package domain

type Piece interface {
	GetColor() Color
	Move(x int, y int) (newX int, newY int)
}

type Color string

const (
	White Color = "white"
	Black Color = "black"
)
