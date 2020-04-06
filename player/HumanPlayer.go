package player

// HumanPlayer ... a human controlled player
type HumanPlayer struct {
	Name        string
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

// GetStats ... get the stats
func (p HumanPlayer) GetStats() Stats {
	return p.PlayerStats
}

// GetInventory ... get the player's inventory
func (p *HumanPlayer) GetInventory() Inventory {
	return p.Backpack
}
