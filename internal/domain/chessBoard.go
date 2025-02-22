package domain

type ChessBoard struct {
	Board [][]*Square `json:"board"`
}

func NewChessBoard() *ChessBoard {
	board := make([][]*Square, 8)
	for i := range board {
		board[i] = make([]*Square, 8) // Cada fila tiene 8 columnas
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = &Square{X: i, Y: j}
		}
	}

	return &ChessBoard{
		Board: board,
	}
}
