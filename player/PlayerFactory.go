package player

func GetPlayer(playerType, name string) Player {
	switch playerType {
	case "HUMAN":
		newPlayer := HumanPlayer{Name: name}
		return newPlayer

	case "COMPUTER":
		newPlayer := ComputerPlayer{Name: name}
		return newPlayer
	default:
		newPlayer := ComputerPlayer{}
		return newPlayer
	}
}
