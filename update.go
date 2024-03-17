package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	g.fps = (g.fps + 1) % 60 // "FPS"
	g.updatePlayer()
	return nil
}

func (g *Game) updatePlayer() {
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
	// ypos
	if g.pYpos > lastCell {
		g.pYpos = 0
	} else if g.pYpos < 0 {
		g.pYpos = lastCell
	}
}
