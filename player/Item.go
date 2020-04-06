package player

// Item ... an item that can be put in an inventory
type Item interface {
	GetName() string
	Weapon
	Armor
}
