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
const numberOfCells = 20

const cellSize = borderSize + cellLength
const screenSize = cellSize*numberOfCells + borderSize

// Just the number of cells minus 1 for convience
const lastCell = numberOfCells - 1

type Game struct {
	// Player coordinates
	pDir       int
	pSpd       int
	snakeTiles linkedlist.LinkedList[[2]int] // [X, Y] positions

	// Food coordinates
	foodXpos int
	foodYpos int

	// FPS not needed but just for fun (also wanted to keep positions whole numbers)
	fps           int
	emptyMapTiles linkedlist.LinkedList[int]
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Snake Go")

	// Create map
	m := linkedlist.LinkedList[int]{}
	for i := 1; i <= numberOfCells*numberOfCells; i++ {
		m.Prepend(i)
	}

	// Initalize player
	startingXpos := rand.Intn(lastCell)
	startingYpos := rand.Intn(lastCell)
	snakeStart := linkedlist.LinkedList[[2]int]{}
	snakeStart.Prepend([2]int{startingXpos, startingYpos})
	m.RemoveFirstValueMatch(startingXpos + startingYpos)

	// Spawn food
	mapPos, ok := m.GetNodeOnIndex(rand.Intn(m.Size()))
	if !ok {
		panic("Error spawning food!")
	}

	foodXpos := (mapPos.Value / numberOfCells) - 1
	foodYpos := (mapPos.Value % numberOfCells) - 1

	// Run game
	if err := ebiten.RunGame(&Game{pSpd: 10, snakeTiles: snakeStart, emptyMapTiles: m, foodXpos: foodXpos, foodYpos: foodYpos}); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSize, screenSize // Equal l+w screen
}
