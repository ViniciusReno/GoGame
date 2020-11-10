package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	err        error
	background *ebiten.Image
)

// Run this code once at startup
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

// Update the screen
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	if err := screen.DrawImage(background, op); err != nil {
		log.Fatal(err)
	}

	return nil
}

// Main loop
func main() {
	if err := ebiten.Run(update, 640, 480, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
