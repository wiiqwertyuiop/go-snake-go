package main

import (
	linkedlist "gosnakego/utils"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	g.fps = (g.fps + 1) % 60 // "FPS"

	// Handle different game states
	if g.snakeTiles.Size() == 0 {
		g.initalizeGame()
	} else if g.gameOver {
		// TODO: Add reset here
		return nil
	} else {
		g.mainGame()
	}

	return nil
}

func (g *Game) initalizeGame() {
	// Create map
	for i := 0; i < numberOfCells; i++ {
		for j := 0; j < numberOfCells; j++ {
			g.emptyMapTiles.Prepend([2]int{i, j})
		}
	}

	// Initalize player
	startingXpos := rand.Intn(lastCell)
	startingYpos := rand.Intn(lastCell)
	g.snakeTiles.Prepend([2]int{startingXpos, startingYpos})
	g.emptyMapTiles.RemoveFirstValueMatch([2]int{startingXpos, startingYpos})

	// Spawn food
	g.SpawnFood()

	g.gameSpd = 20
}

func (g *Game) mainGame() {
	// Handle input
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && (g.lastDir != 2 || g.snakeTiles.Size() == 1) {
		g.pDir = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && (g.lastDir != 1 || g.snakeTiles.Size() == 1) {
		g.pDir = 2
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) && (g.lastDir != 4 || g.snakeTiles.Size() == 1) {
		g.pDir = 3
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) && (g.lastDir != 3 || g.snakeTiles.Size() == 1) {
		g.pDir = 4
	}

	// Update every certain number of frames
	if g.fps%g.gameSpd == 0 && g.pDir != 0 {
		g.updateMap(g.updatePlayer())
	}
}

func (g *Game) updatePlayer() [2]int {
	curPos, ok := g.snakeTiles.ReadValueOnIndex(0)
	if !ok {
		panic("Problem moving the snake!")
	}

	// Move player a direction
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
	g.lastDir = g.pDir

	// Loop player if off screen
	if curPos[0] > lastCell {
		curPos[0] = 0
	} else if curPos[0] < 0 {
		curPos[0] = lastCell
	}
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
	g.emptyMapTiles.RemoveFirstValueMatch(newPos)
	head := true
	// Check for snake contact, and remove old snake pieces
	g.snakeTiles.LoopActionOnListRemoveNextOnFalse(func(curNode *linkedlist.Node[[2]int]) bool {
		// Handle snake touching itself
		if newPos == curNode.Value {
			if !head {
				g.pDir = 0
				g.gameOver = true
			}
			head = false
		}
		// Handle tail
		if _, moreSnake := curNode.Offset(2); !moreSnake {
			oldTail, ok := curNode.Offset(1)
			if !ok {
				panic("Something went wrong removing the tail")
			}
			if curNode.Value != oldTail.Value {
				// The old tail is now a valid empty space
				g.emptyMapTiles.Prepend(oldTail.Value)
			}
			return false
		}
		return true
	})

	// Check for food
	if newPos[0] == g.foodXpos && newPos[1] == g.foodYpos {
		g.snakeTiles.Prepend(newPos)
		g.snakeTiles.Prepend(newPos)
		g.SpawnFood()
		g.score++

		// Increase speed every ertain amount of food collected
		if g.score%5 == 0 && g.gameSpd > 8 {
			g.gameSpd--
		}
	}
}

func (g *Game) SpawnFood() {
	mapPos, ok := g.emptyMapTiles.GetNodeOnIndex(rand.Intn(g.emptyMapTiles.Size()))
	if !ok {
		panic("Error spawning food!")
	}

	g.foodXpos = mapPos.Value[0]
	g.foodYpos = mapPos.Value[1]
}
