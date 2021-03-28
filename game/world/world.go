package world

import (
	"github.com/edznux/codagotchi/game/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	Height int
	Width  int
	Speed  int
	Name   string
	Seed   int64
	Tick   int64
}

func (w *World) Draw(screen *ebiten.Image) {
}

func (w *World) Init() {
}

func (w *World) GetRandomPosInBound() (float64, float64) {
	posX := rand.Float64(w.Tick) * float64(w.Width)
	// Need 2 distinct seed to get 2 different values.
	posY := rand.Float64(w.Tick+1) * float64(w.Height)
	return posX, posY
}

func (w *World) Update() {
	w.Tick++
}
