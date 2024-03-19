package main

import (
	linkedlist "gosnakego/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	g.fps = (g.fps + 1) % 60 // "FPS"
	g.updateMap(g.updatePlayer())
	return nil
}

func (g *Game) updatePlayer() [2]int {
	curPos, ok := g.snakeTiles.ReadValueOnIndex(0)
	if !ok {
		panic("Problem moving the snake!")
	}

	// Move player
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && (g.pDir != 2 || g.snakeTiles.Size() == 1) {
		g.pDir = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && (g.pDir != 1 || g.snakeTiles.Size() == 1) {
		g.pDir = 2
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) && (g.pDir != 4 || g.snakeTiles.Size() == 1) {
		g.pDir = 3
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) && (g.pDir != 3 || g.snakeTiles.Size() == 1) {
		g.pDir = 4
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		// Debug: increase snake size
		for i := 0; i < 10; i++ {
			g.snakeTiles.Prepend(curPos)
		}
	}

	// Update speed every certain number of frames
	if g.fps%g.pSpd == 0 {
		switch g.pDir {
		case 1:
			curPos[0]++
		case 2:
			curPos[0]--
		case 3:
			curPos[1]--
		case 4:
			curPos[1]++
		}
	}

	// Loop player if off screen
	if curPos[0] > lastCell {
		curPos[0] = 0
	} else if curPos[0] < 0 {
		curPos[0] = lastCell
	}
	// ypos
	if curPos[1] > lastCell {
		curPos[1] = 0
	} else if curPos[1] < 0 {
		curPos[1] = lastCell
	}
	return curPos
}

func (g *Game) updateMap(newPos [2]int) {
	// Update snake on map
	g.snakeTiles.Prepend(newPos)
	delete(g.emptyMapTiles, newPos[0]+newPos[1])

	// Check for snake contact, and remove old snake pieces
	g.snakeTiles.LoopActionOnListRemoveNextOnFalse(func(curNode *linkedlist.Node[[2]int]) bool {
		// Handle tail
		if _, moreSnake := curNode.Offset(2); !moreSnake {
			oldTail, ok := curNode.Offset(1)
			if !ok {
				panic("Something went wrong removing the tail")
			}
			// The old tail is now a valid empty space
			g.emptyMapTiles[oldTail.Value[0]+oldTail.Value[1]] = true
			return false
		}
		return true
	})
}
