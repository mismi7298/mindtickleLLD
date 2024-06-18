package board

import (
	"chess/pkg/piece"
)

type Board struct {
	Piece [][]piece.Piece
}

func NewBoard() *Board {
	return &Board{}
}

type BoardInterface interface {
}
