package main

import (
	linkedlist "gosnakego/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game setup
const cellLength = 10
const borderSize = 0
const numberOfCells = 20

const cellSize = borderSize + cellLength
const screenSize = cellSize*numberOfCells + borderSize

// Just the number of cells minus 1 for convience
const lastCell = numberOfCells - 1

type Game struct {
	// Player
	pDir       int
	lastDir    int
	score      int
	snakeTiles linkedlist.LinkedList[[2]int] // [X, Y] positions

	// Food coordinates
	foodXpos int
	foodYpos int

	// frameCount not needed, but just for fun (also wanted to keep positions whole numbers)
	frameCount    int
	updateFrame   int
	emptyMapTiles linkedlist.LinkedList[[2]int]

	gameOver bool
}

func main() {
	initFont()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Snake Go")

	// Run game
	if err := ebiten.RunGame(&Game{updateFrame: 15}); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSize, screenSize // Equal l+w screen
}
