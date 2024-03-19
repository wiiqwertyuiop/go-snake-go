package main

import (
	linkedlist "gosnakego/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen
	screen.Fill(color.RGBA{15, 15, 15, 0})

	// Draw Grid
	for i := 1; i <= screenSize; i += cellSize {
		vector.StrokeLine(screen, float32(i), 0, float32(i), float32(screenSize), float32(borderSize), color.White, false)
		vector.StrokeLine(screen, 0, float32(i), float32(screenSize), float32(i), float32(borderSize), color.White, false)
	}

	// Draw player
	g.snakeTiles.LoopActionOnListRemoveNextOnFalse(func(curNode *linkedlist.Node[[2]int]) bool {
		tile := curNode.Value
		vector.DrawFilledRect(screen, float32(tile[0]*cellSize+borderSize), float32(tile[1]*cellSize+borderSize), 10, 10, color.RGBA{150, 0, 0, 0}, false)
		return true
	})

	// Draw food
	vector.DrawFilledRect(screen, float32(g.foodXpos*cellSize+borderSize), float32(g.foodYpos*cellSize+borderSize), 10, 10, color.RGBA{0, 150, 0, 0}, false)

}
