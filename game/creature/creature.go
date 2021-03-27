package creature

import (
	"bytes"
	"image"
	"log"
	"math"

	"github.com/edznux/codagotchi/game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Creature struct {
	Name string

	Life            int64
	LifeSpanCounter int64 //increment every tick

	PosX float64
	PosY float64

	VelocityX float64
	VelocityY float64

	TargetX float64
	TargetY float64

	Speed float64

	Image *ebiten.Image `json:"-"`
}

func (c *Creature) GetCenterPos() (x float64, y float64) {
	x = c.PosX - float64(c.Image.Bounds().Dx()/2)
	y = c.PosY - float64(c.Image.Bounds().Dy()/2)
	return x, y
}

func (c *Creature) DistanceToPoint(toX, toY float64) float64 {
	todoX := toX - c.PosX
	todoY := toY - c.PosY

	todo := math.Sqrt(math.Pow(todoX, 2) + math.Pow(todoY, 2))
	return todo
}

func (c *Creature) MoveTo(toX, toY float64) {
	distance := c.DistanceToPoint(toX, toY)
	c.VelocityX = (toX - c.PosX) / distance
	c.VelocityY = (toY - c.PosY) / distance

	// if the create is close enough to the point, teleport it to the destination
	// and remove any velocity. This avoid the small 1px jitter back and forth.
	if distance < 1 {
		c.PosX = toX
		c.PosY = toY

		c.VelocityX = 0
		c.VelocityY = 0
	}
}

func (c *Creature) Update() {
	c.LifeSpanCounter++

	c.MoveTo(c.TargetX, c.TargetY)
	c.PosX += c.Speed * c.VelocityX
	c.PosY += c.Speed * c.VelocityY

	// c.PosX = c.PosX + (c.VelocityX * c.PosX)
	// c.PosY = c.PosY + (c.VelocityY * c.PosY)
}

func (c *Creature) Init() {
	petPng, _, err := image.Decode(bytes.NewReader(assets.PetV1_png))
	if err != nil {
		log.Fatal(err)
	}
	bobImg := ebiten.NewImageFromImage(petPng)
	c.Image = bobImg
}

func (c *Creature) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.GetCenterPos())
	screen.DrawImage(c.Image, op)
}
