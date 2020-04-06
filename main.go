package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kvitebjorn/umm/player"
)

// This is required to draw debug texts.
func update(screen *ebiten.Image) error {
	luka := player.GetPlayer("HUMAN", "Luka")
	ebitenutil.DebugPrint(screen, "Hello, "+luka.GetName())
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Umm"); err != nil {
		panic(err)
	}
}
