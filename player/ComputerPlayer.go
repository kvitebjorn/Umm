package player

// ComputerPlayer ... a computer controlled player, usually a monster
type ComputerPlayer struct {
	Name        string
	Whereabouts Position
	PlayerStats Stats
}

// GetName ... get the player name
func (p ComputerPlayer) GetName() string {
	return p.Name
}

// GetPosition ... get the player position
func (p ComputerPlayer) GetPosition() Position {
	return p.Whereabouts
}

// GetStats ... get the stats
func (p ComputerPlayer) GetStats() Stats {
	return p.PlayerStats
}
