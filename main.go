package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game setup
const cellLength = 10
const borderSize = 0
const numberOfCells = 50

const cellSize = borderSize + cellLength
const screenSize = cellSize*numberOfCells + borderSize

// Just the number of cells minus 1 for convience
const lastCell = numberOfCells - 1

type Game struct {
	// Player coordinates
	pXpos int
	pYpos int
	pDir  int
	pSpd  int

	// Not needed but just for fun (also wanted to keep positions whole numbers)
	fps int
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Snake Go")
	if err := ebiten.RunGame(&Game{pXpos: rand.Intn(lastCell), pYpos: rand.Intn(lastCell), pSpd: 3}); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSize, screenSize // Equal l+w screen
}
