package main

import (
	linkedlist "gosnakego/utils"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game setup
const cellLength = 10
const borderSize = 0
const numberOfCells = 35

const cellSize = borderSize + cellLength
const screenSize = cellSize*numberOfCells + borderSize

// Just the number of cells minus 1 for convience
const lastCell = numberOfCells - 1

type Game struct {
	// Player coordinates
	pDir       int
	pSpd       int
	snakeTiles linkedlist.LinkedList[[2]int] // [X, Y] positions

	// FPS not needed but just for fun (also wanted to keep positions whole numbers)
	fps           int
	emptyMapTiles map[int]bool
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Snake Go")

	// Create map
	m := map[int]bool{}
	for i := 0; i < numberOfCells*numberOfCells; i++ {
		m[i] = true
	}

	// Initalize player
	startingXpos := rand.Intn(lastCell)
	startingYpos := rand.Intn(lastCell)
	snakeStart := linkedlist.LinkedList[[2]int]{}
	snakeStart.Prepend([2]int{startingXpos, startingYpos})
	delete(m, startingXpos+startingYpos)

	// Run game
	if err := ebiten.RunGame(&Game{pSpd: 3, snakeTiles: snakeStart, emptyMapTiles: m}); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSize, screenSize // Equal l+w screen
}
