package creature

import "github.com/hajimehoshi/ebiten/v2"

type Creature struct {
	Name string

	Life            int64
	LifeSpanCounter int64 //increment every tick

	PosX float64
	PosY float64

	VelocityX float64
	VelocityY float64

	Image *ebiten.Image `json:"-"`
}

func (c *Creature) GetCenterPos() (x float64, y float64) {
	x = c.PosX - float64(c.Image.Bounds().Dx()/2)
	y = c.PosY - float64(c.Image.Bounds().Dy()/2)
	return x, y
}

func (c *Creature) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.GetCenterPos())
	screen.DrawImage(c.Image, op)
}
