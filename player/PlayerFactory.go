package player

// GetPlayer ... create a Player instance
func GetPlayer(playerType, name string) Player {
	switch playerType {
	case "HUMAN":
		newPlayer := HumanPlayer{Name: name}
		newPlayerStats := Stats{
			Hp:  Stat{Name: "HP", Xp: 1, CurrentLevel: 10},
			Str: Stat{Name: "Strength", Xp: 1, CurrentLevel: 1},
			Def: Stat{Name: "Defense", Xp: 1, CurrentLevel: 1},
			Atk: Stat{Name: "Attack", Xp: 1, CurrentLevel: 1},
		}
		newPlayer.PlayerStats = newPlayerStats
		newPlayer.Helmet = BronzeHelmet{}
		newPlayer.Chestplate = BronzeChest{}
		newPlayer.Legplate = BronzeLeg{}
		newPlayer.Sword = BronzeSword{}
		newPlayer.Backpack = Inventory{}
		return newPlayer

	case "COMPUTER":
		newPlayer := ComputerPlayer{Name: name}
		return newPlayer
	default:
		newPlayer := ComputerPlayer{}
		newPlayerStats := Stats{
			Hp:  Stat{Name: "HP", Xp: 1, CurrentLevel: 1},
			Str: Stat{Name: "Strength", Xp: 1, CurrentLevel: 1},
			Def: Stat{Name: "Defense", Xp: 1, CurrentLevel: 1},
			Atk: Stat{Name: "Attack", Xp: 1, CurrentLevel: 1},
		}
		newPlayer.PlayerStats = newPlayerStats
		return newPlayer
	}
}
