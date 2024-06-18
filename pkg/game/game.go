package game

import (
	"chess/pkg/board"
	"chess/pkg/location"
	"chess/pkg/piece"
	"chess/pkg/player"
)

type Game struct {
	Players map[int]player.Player
	Board   board.Board
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) SetPlayer(id int, player player.Player) {
	g.Players[id] = player
}

type GameInterface interface {
	IsEnded() bool
	Play()
	MakeMove(player.Player, piece.Piece, location.Location)
}
