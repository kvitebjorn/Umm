package player

// BronzeBaseDefensePower ... the base defense power for bronze armor
const BronzeBaseDefensePower = 1

// Armor ... will have attack powers, etc
type Armor interface {
	HeadArmor
	ChestArmor
	LegArmor
}

// GetBronzeBaseDefensePower ... returns bronze base defense multiplier
func GetBronzeBaseDefensePower() int {
	return BronzeBaseDefensePower
}
