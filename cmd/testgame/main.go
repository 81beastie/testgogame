package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/test-go-game/game-hello-world/internal/assets"
	"github.com/test-go-game/game-hello-world/internal/controls"
	"github.com/test-go-game/game-hello-world/internal/game"

	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
)

func main() {
	g := &myGame{
		ctx: &game.Context{
			WindowWidth:  320,
			WindowHeight: 240,
		},
	}

	g.ctx.Loader = createLoader()

	g.Init()

	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.ctx.Input = g.inputSystem.NewHandler(0, controls.DefaultKeymap)

	ebiten.SetWindowSize(g.ctx.WindowWidth, g.ctx.WindowWidth)
	ebiten.SetWindowTitle("HW")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type myGame struct {
	//windowWidth  int
	//windowHeight int
	//loader       *resource.Loader
	player      *Player
	inputSystem input.System
	//input        *input.Handler
	ctx *game.Context
}

func (g *myGame) Init() {
	assets.RegisterResources(g.ctx.Loader)
	gopher := g.ctx.Loader.LoadImage(assets.ImageGopher).Data
	g.player = &Player{img: gopher}
}

func (g *myGame) Update() error {
	g.inputSystem.Update()

	speed := 64.0 * (1.0 / 60)
	var v gmath.Vec

	if g.ctx.Input.ActionIsPressed(controls.ActionMoveRight) {
		v.X += speed
	}

	if g.ctx.Input.ActionIsPressed(controls.ActionMoveLeft) {
		v.X -= speed
	}

	if g.ctx.Input.ActionIsPressed(controls.ActionMoveUp) {
		v.Y -= speed
	}

	if g.ctx.Input.ActionIsPressed(controls.ActionMoveDown) {
		v.Y += speed
	}

	g.player.pos = g.player.pos.Add(v)
	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {

	gopher := g.ctx.Loader.LoadImage(assets.ImageGopher).Data
	var options ebiten.DrawImageOptions
	options.GeoM.Translate(g.player.pos.X, g.player.pos.Y)
	screen.DrawImage(gopher, &options)
}

func (g *myGame) Layout(w, h int) (int, int) {

	return g.ctx.WindowWidth, g.ctx.WindowHeight
}

func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}

type Player struct {
	pos gmath.Vec // {X, Y}
	img *ebiten.Image
}
