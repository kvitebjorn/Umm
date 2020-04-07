package player

// Position ... the player position in the world grid
type Position struct {
	X int
	Y int
}

// Player ... stuff that all humans & computer players should do
type Player interface {
	GetName() string
	GetStats() Stats
	GetPosition() Position
	SetPosition(pos Position)
}
