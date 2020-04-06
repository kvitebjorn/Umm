package player

// ChestArmorBaseDefensePower ... the base defense power for chest armor
const ChestArmorBaseDefensePower = 3

// ChestArmor ... armor for the chest slot
type ChestArmor interface {
	GetChestDefensePower() int
}

// BronzeChest ... a chest made of bronze metal
type BronzeChest struct {
	DefensePower int
}

// TODO: make this generic to the type of metal in the future, don't write one of these out for each metal type

// GetChestDefensePower ... get the defense power of a bronze chest
func (b *BronzeChest) GetChestDefensePower() int {
	b.DefensePower = GetBronzeBaseDefensePower() * ChestArmorBaseDefensePower
	return b.DefensePower
}

// GetName ... the name of the bronze chest
func (b *BronzeChest) GetName() string {
	return "Bronze Chestplate"
}
