package world

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

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

func (w *World) Init() {
}

func (w *World) GetRandomPosInBound() (float64, float64) {
	posX := rand.Float64() * float64(w.Width)
	posY := rand.Float64() * float64(w.Height)
	return posX, posY
}

func (w *World) Update() {
	w.Tick++
}
