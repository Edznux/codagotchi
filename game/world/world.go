package world

import "github.com/hajimehoshi/ebiten/v2"

type World struct {
	Height int
	Width  int
	Speed  int
	Name   string
	Seed   int
	Tick   int64
}

func (w *World) Draw(screen *ebiten.Image) {
}
