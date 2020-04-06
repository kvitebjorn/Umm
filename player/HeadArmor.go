package player

// HeadArmorBaseDefensePower ... the base defense power for head armor
const HeadArmorBaseDefensePower = 1

// HeadArmor ... armor for head slot
type HeadArmor interface {
	GetHeadDefensePower() int
}

// BronzeHelmet ... a helmet made of bronze metal
type BronzeHelmet struct {
	DefensePower int
}

// TODO: make this generic to the type of metal in the future, don't write one of these out for each metal type

// GetHeadDefensePower ... get the defense power of a bronze helmet
func (b *BronzeHelmet) GetHeadDefensePower() int {
	b.DefensePower = GetBronzeBaseDefensePower() * HeadArmorBaseDefensePower
	return b.DefensePower
}

// GetName ... the name of the bronze helmet
func (b *BronzeHelmet) GetName() string {
	return "Bronze Helmet"
}
