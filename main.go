package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const cellLength = 10
const borderSize = 1
const numberOfCells = 10

const screenSize = (borderSize+cellLength)*numberOfCells + borderSize

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw Grid
	for i := 1; i <= screenSize; i += (cellLength + borderSize) {
		vector.StrokeLine(screen, float32(i), 0, float32(i), float32(screenSize), float32(borderSize), color.White, false)
		vector.StrokeLine(screen, 0, float32(i), float32(screenSize), float32(i), float32(borderSize), color.White, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSize, screenSize // Square screen
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Snake Go")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
