package player

// LegArmorBaseDefensePower ... the base defense power for leg armor
const LegArmorBaseDefensePower = 2

// LegArmor ... armor for the leg slot
type LegArmor interface {
	GetLegDefensePower() int
}

// BronzeLeg ... a leg made of bronze metal
type BronzeLeg struct {
	DefensePower int
}

// TODO: make this generic to the type of metal in the future, don't write one of these out for each metal type

// GetLegDefensePower ... get the defense power of a bronze leg
func (b *BronzeLeg) GetLegDefensePower() int {
	b.DefensePower = GetBronzeBaseDefensePower() * LegArmorBaseDefensePower
	return b.DefensePower
}

// GetName ... the name of the bronze legs
func (b *BronzeLeg) GetName() string {
	return "Bronze Legplate"
}
