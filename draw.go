package main

import (
	linkedlist "gosnakego/utils"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusBigFont font.Face
)

func initFont() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    numberOfCells,
		DPI:     72,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen
	screen.Fill(color.RGBA{15, 15, 15, 0})

	// Draw Grid
	if borderSize != 0 {
		for i := 1; i <= screenSize; i += cellSize {
			vector.StrokeLine(screen, float32(i), 0, float32(i), float32(screenSize), float32(borderSize), color.White, false)
			vector.StrokeLine(screen, 0, float32(i), float32(screenSize), float32(i), float32(borderSize), color.White, false)
		}
	}

	// Handle game overs
	playerColor := color.RGBA{150, 0, 0, 0}
	if g.gameOver {
		playerColor = color.RGBA{150, 0, 150, 0}
		text.Draw(screen, "GAME OVER", mplusBigFont, numberOfCells+numberOfCells, numberOfCells+numberOfCells, color.White)
	}

	// Draw player
	g.snakeTiles.LoopActionOnListRemoveNextOnFalse(func(curNode *linkedlist.Node[[2]int]) bool {
		tile := curNode.Value
		vector.DrawFilledRect(screen, float32(tile[0]*cellSize+borderSize), float32(tile[1]*cellSize+borderSize), 10, 10, playerColor, false)
		return true
	})

	// Draw food
	vector.DrawFilledRect(screen, float32(g.foodXpos*cellSize+borderSize), float32(g.foodYpos*cellSize+borderSize), 10, 10, color.RGBA{0, 150, 0, 0}, false)
}
