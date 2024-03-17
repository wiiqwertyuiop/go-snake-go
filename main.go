package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func (g *Game) Update() error {
	g.fps = (g.fps + 1) % 60 // "FPS"

	// Move player
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.pDir = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.pDir = 2
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.pDir = 3
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.pDir = 4
	}

	// Update speed every certain number of frames
	if g.fps%g.pSpd == 0 {
		switch g.pDir {
		case 1:
			g.pXpos++
		case 2:
			g.pXpos--
		case 3:
			g.pYpos--
		case 4:
			g.pYpos++
		}
	}

	// Loop player if off screen
	if g.pXpos > lastCell {
		g.pXpos = 0
	} else if g.pXpos < 0 {
		g.pXpos = lastCell
	}
	if g.pYpos > lastCell {
		g.pYpos = 0
	} else if g.pYpos < 0 {
		g.pYpos = lastCell
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen
	screen.Fill(color.RGBA{15, 15, 15, 0})

	// Draw Grid
	for i := 1; i <= screenSize; i += cellSize {
		vector.StrokeLine(screen, float32(i), 0, float32(i), float32(screenSize), float32(borderSize), color.White, false)
		vector.StrokeLine(screen, 0, float32(i), float32(screenSize), float32(i), float32(borderSize), color.White, false)
	}
	// Draw player
	vector.DrawFilledRect(screen, float32(g.pXpos*cellSize+borderSize), float32(g.pYpos*cellSize+borderSize), 10, 10, color.RGBA{150, 0, 0, 0}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSize, screenSize // Equal l+w screen
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Snake Go")
	if err := ebiten.RunGame(&Game{pXpos: rand.Intn(lastCell), pYpos: rand.Intn(lastCell), pSpd: 3}); err != nil {
		log.Fatal(err)
	}
}
