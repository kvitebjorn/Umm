package player

// BronzeBaseAttackPower ... the base attack power of bronze weapons
const BronzeBaseAttackPower = 1

// Weapon ... will have attack powers, etc
type Weapon interface {
	SwordWeapon
}

// GetBronzeBaseAttackPower ... returns bronze base attack multiplier
func GetBronzeBaseAttackPower() int {
	return BronzeBaseAttackPower
}
