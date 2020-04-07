package player

// HumanPlayer ... a human controlled player
type HumanPlayer struct {
	Name        string
	Whereabouts Position
	PlayerStats Stats
	Backpack    Inventory

	Helmet     HeadArmor
	Chestplate ChestArmor
	Legplate   LegArmor
	Sword      SwordWeapon
}

// GetName ... get the player name
func (p HumanPlayer) GetName() string {
	return p.Name
}

// GetPosition ... get the player's position
func (p HumanPlayer) GetPosition() Position {
	return p.Whereabouts
}

// GetStats ... get the stats
func (p HumanPlayer) GetStats() Stats {
	return p.PlayerStats
}

// GetInventory ... get the player's inventory
func (p *HumanPlayer) GetInventory() Inventory {
	return p.Backpack
}
