package world

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	Height  int
	Width   int
	Speed   int
	Name    string
	Seed    int64
	randGen *rand.Rand
	Tick    int64
}

func (w *World) Draw(screen *ebiten.Image) {
}

func (w *World) Init() {
	w.randGen = rand.New(rand.NewSource(w.Seed))
}

func (w *World) GetRandomPosInBound() (float64, float64) {
	posX := w.randGen.Float64() * float64(w.Width)
	posY := w.randGen.Float64() * float64(w.Height)
	return posX, posY
}

func (w *World) Update() {
	w.Tick++
}
