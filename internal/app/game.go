package app

import (
	domain "chess/internal/domain"
	"chess/internal/domain/pieces"
)

type Game struct {
	ChessBoard *domain.ChessBoard
}

func NewGame() *Game {
	game := &Game{
		ChessBoard: domain.NewChessBoard(),
	}
	game.fillBoard(game.ChessBoard)

	return game
}

func (g *Game) fillBoard(board *domain.ChessBoard) {

	var piece domain.Piece = pieces.NewPawn(domain.White)
	board.Board[0][0].Piece = &piece
	board.Board[0][1].Piece = &piece
	board.Board[0][2].Piece = &piece
	board.Board[0][3].Piece = &piece
	board.Board[0][4].Piece = &piece
	board.Board[0][5].Piece = &piece

}
