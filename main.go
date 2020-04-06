package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/kvitebjorn/umm/player"
)

const (
	screenWidth  = 240
	screenHeight = 240
)

const (
	tileSize = 16
	tileXNum = 25
)

var (
	tilesImage *ebiten.Image
)

func init() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

var (
	text            = ""
	counter         = 0
	playerCreated   = false
	continuePressed = false
	player1         = player.GetPlayer("HUMAN", "")

	layers = [][]int{
		{
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 244, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 219, 243, 243, 243, 219, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243,
			243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
		},
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 51, 52, 53, 54, 55, 56, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 76, 77, 78, 79, 80, 81, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 101, 102, 103, 104, 105, 106, 0, 0, 0, 0,

			0, 0, 0, 0, 0, 126, 127, 128, 129, 130, 131, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 303, 303, 245, 242, 303, 303, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,

			0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
		},
	}
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

	return nil
}

// Updates the current frame (60 frames occur per second)
func update(screen *ebiten.Image) error {

	if !continuePressed {
		doPlayerLogin(screen)
		return nil
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw each tile with each DrawImage call.
	// As the source images of all DrawImage calls are always same,
	// this rendering is done very effectively.
	// For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten#Image.DrawImage
	const xNum = screenWidth / tileSize
	for _, l := range layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

			sx := (t % tileXNum) * tileSize
			sy := (t / tileXNum) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Umm"); err != nil {
		panic(err)
	}
}
