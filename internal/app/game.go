package app

import (
	domain "chess/internal/domain"
	"chess/internal/domain/pieces"
	"log"
)

type Game struct {
	ChessBoard *domain.ChessBoard
}

func NewGame() *Game {
	game := &Game{
		ChessBoard: domain.NewChessBoard(),
	}
	game.fillBoard(game.ChessBoard, 0, domain.White)
	game.fillBoard(game.ChessBoard, 7, domain.Black)

	return game
}

func (g *Game) fillBoard(board *domain.ChessBoard, row int, color domain.Color) {

	var rook domain.Piece = pieces.NewRook(color)
	board.Board[row][0].Piece = &rook
	board.Board[row][7].Piece = &rook

	var knight domain.Piece = pieces.NewKnight(color)
	board.Board[row][1].Piece = &knight
	board.Board[row][6].Piece = &knight

	var bishop domain.Piece = pieces.NewBishop(color)
	board.Board[row][2].Piece = &bishop
	board.Board[row][5].Piece = &bishop

	var king domain.Piece = pieces.NewKing(color)
	board.Board[row][3].Piece = &king

	var queen domain.Piece = pieces.NewQueen(color)
	board.Board[row][4].Piece = &queen

	if color == domain.White {
		row++
	} else {
		row--
	}

	var pawn domain.Piece = pieces.NewPawn(color)
	board.Board[row][0].Piece = &pawn
	board.Board[row][1].Piece = &pawn
	board.Board[row][2].Piece = &pawn
	board.Board[row][3].Piece = &pawn
	board.Board[row][4].Piece = &pawn
	board.Board[row][5].Piece = &pawn
	board.Board[row][6].Piece = &pawn
	board.Board[row][7].Piece = &pawn

}

func (g *Game) MovePiece(x1, y1, x2, y2 int) {

	// Verificar si las coordenadas están dentro de los límites del tablero
	if x1 < 0 || x1 >= 8 || y1 < 0 || y1 >= 8 || x2 < 0 || x2 >= 8 || y2 < 0 || y2 >= 8 {
		return
	}

	// Verificar si hay una pieza en la posición de origen
	piece := g.ChessBoard.Board[x1][y1].Piece
	piece2 := g.ChessBoard.Board[x2][y2].Piece
	isVacia := true

	if piece == nil {
		log.Print("la pieza es nula")
		return
	}

	if piece2 != nil {
		isVacia = false
		if (*piece).GetColor() == (*piece2).GetColor() {
			log.Print("la pieza es del mismo color")
			return
		}
	}

	if (*piece).Move(x1, y1, x2, y2, isVacia) {

		g.ChessBoard.Board[x2][y2].Piece = piece
		g.ChessBoard.Board[x1][y1].Piece = nil

	} else {
		log.Print("Movimiento   raro válido")
		return
	}

}
