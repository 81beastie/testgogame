package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	g := &myGame{
		windowWidth:  320,
		windowHeight: 240,
	}

	ebiten.SetWindowSize(g.windowWidth, g.windowWidth)
	ebiten.SetWindowTitle("HW")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type myGame struct {
	windowWidth  int
	windowHeight int
}

func (g *myGame) Update() error {
	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *myGame) Layout(w, h int) (int, int) {
	// Layout - тема для продвинутых, поэтому нам пока
	// достаточно считать, что screen size = window size.
	return g.windowWidth, g.windowHeight
}
