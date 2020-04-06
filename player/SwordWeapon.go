package player

// SwordWeaponBaseAttackPower ... the base attack power for swords
const SwordWeaponBaseAttackPower = 1

// SwordWeapon ... a weapon of type sword
type SwordWeapon interface {
	GetSwordAttackPower() int
}

// BronzeSword ... a sword made of bronze metal
type BronzeSword struct {
	AttackPower int
}

// TODO: make this generic to the type of metal in the future, don't write one of these out for each metal type

// GetSwordAttackPower ... get the attack power of a bronze sword
func (b *BronzeSword) GetSwordAttackPower() int {
	b.AttackPower = GetBronzeBaseAttackPower() * SwordWeaponBaseAttackPower
	return b.AttackPower
}

// GetName ... the name of the bronze sword
func (b *BronzeSword) GetName() string {
	return "Bronze Sword"
}
