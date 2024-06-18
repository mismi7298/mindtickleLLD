package piece

import "chess/pkg/location"

type Status int

const (
	InGame Status = iota + 1
	OutOfGame
)

type Piece struct {
	position location.Location
	status   Status
}

func NewPiece() *Piece {
	return &Piece{}
}

type PieceInterface interface {
	ValidMove(newLocation location.Location) bool
	UpdatePiece(location.Location, Status)
}
