package player

// Player ... stuff that all humans & computer players should do
type Player interface {
	GetName() string
	GetStats() Stats
}
