package main

import (
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/kvitebjorn/umm/player"
)

var (
	text            = ""
	counter         = 0
	playerCreated   = false
	continuePressed = false
	player1         = player.GetPlayer("HUMAN", "")
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

func updatePlayerCreationDialogueState(screen *ebiten.Image) {

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

	// If the enter key is pressed, create a human player
	if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyKPEnter) {
		player1 = player.GetPlayer("HUMAN", text)
		playerCreated = true
	}

	return
}

func updateContinueDialogueState(screen *ebiten.Image) {
	// If the enter key is pressed again, continue
	if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyKPEnter) {
		continuePressed = true
	}
}

func showPlayerCreationDialogue(screen *ebiten.Image) {
	t := text
	if counter%60 < 30 {
		t += "_"
	}
	ebitenutil.DebugPrint(screen, "What is your name?\n"+t)
}

func showContinueDialogue(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, "+player1.GetName()+"\nPress ENTER to continue...")
}

func doPlayerLogin(screen *ebiten.Image) error {
	if !playerCreated {
		showPlayerCreationDialogue(screen)
		updatePlayerCreationDialogueState(screen)
		counter = 0
		return nil
	} else if !continuePressed {
		showContinueDialogue(screen)
		counter++
		if counter > 60 {
			updateContinueDialogueState(screen)
		}
		return nil
	}

	if continuePressed {
		ebitenutil.DebugPrint(screen, "Welcome to Umm...")
	}

	return nil
}

// Updates the current frame (60 frames occur per second)
func update(screen *ebiten.Image) error {

	doPlayerLogin(screen)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Umm"); err != nil {
		panic(err)
	}
}
