package player

type Player struct {
	Rank int
}

func NewPlayer() *Player {
	return &Player{}
}

type PlayerInterface interface {
	// TryMove(piece, location )
}
