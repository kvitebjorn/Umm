package player

// ComputerPlayer ... a computer controlled player, usually a monster
type ComputerPlayer struct {
	Name        string
	PlayerStats Stats
}

// GetName ... get the player name
func (p *ComputerPlayer) GetName() string {
	return p.Name
}

// GetStats ... get the stats
func (p *ComputerPlayer) GetStats() Stats {
	return p.PlayerStats
}
