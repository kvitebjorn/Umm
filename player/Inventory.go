package player

// InventoryLimit ... the max # of items the inventory can hold
const InventoryLimit = 25

// Inventory ... a bag of stuff! Limited to 25 items right now
type Inventory struct {
	Items [InventoryLimit]Item
}
