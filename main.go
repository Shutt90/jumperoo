package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct{}

var gopher *ebiten.Image
var grassPlatform *ebiten.Image

func init() {
	var err error
	gopher, _, err = ebitenutil.NewImageFromFile("static/gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	grassPlatform, _, err = ebitenutil.NewImageFromFile("static/grassplatform.png")
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.2, 0.2)
	op.GeoM.Translate(screenWidth/2.25, screenHeight*0.75)

	screen.DrawImage(gopher, op)
	
	screen.DrawImage(grassPlatform, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return 800, 600
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Jumperoo")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
