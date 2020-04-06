package main

import (
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/kvitebjorn/umm/player"
)

var (
	text          = ""
	counter       = 0
	playerCreated = false
	player1       = player.GetPlayer("HUMAN", "")
)

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

// This is required to draw debug texts.
func update(screen *ebiten.Image) error {

	// Add a string from InputChars, that returns string input by users.
	// Note that InputChars result changes every frame, so you need to call this
	// every frame.
	text += string(ebiten.InputChars())

	// Adjust the string to be at most 1 line.
	ss := strings.Split(text, "\n")
	if len(ss) > 10 {
		text = strings.Join(ss[len(ss)-1:], "\n")
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(text) >= 1 {
			text = text[:len(text)-1]
		}
	}

	counter++

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if playerCreated {
		ebitenutil.DebugPrint(screen, "Hello, "+player1.GetName())
	}

	if !playerCreated {

		// If the enter key is pressed, create a human player
		if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyKPEnter) {
			player1 = player.GetPlayer("HUMAN", text)
			playerCreated = true
			return nil
		}

		// Blink the cursor.
		t := text
		if counter%60 < 30 {
			t += "_"
		}
		ebitenutil.DebugPrint(screen, "What is your name?\n"+t)
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Umm"); err != nil {
		panic(err)
	}
}
