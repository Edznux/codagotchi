package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/edznux/codagotchi/game/assets"
	"github.com/edznux/codagotchi/game/creature"
	"github.com/edznux/codagotchi/game/world"
	"github.com/edznux/codagotchi/metrics"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	// No need to save the name of the JSON in the json...
	SaveName string `json:"-"`

	// Bob is our little pet, at least, for now :)
	Bob   *creature.Creature
	World *world.World

	// Don't save the Statsd client in the JSON of the save
	Statsd *statsd.Client `json:"-"`
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.Bob.LifeSpanCounter++
	g.World.Tick++
	// every 10 seconds
	if g.World.Tick%600 == 0 {
		g.Statsd.Gauge("codagotchi.bob.lifespan", float64(g.Bob.LifeSpanCounter), metrics.Tags, 1)
		g.Statsd.Gauge("codagotchi.bob.life", float64(g.Bob.Life), metrics.Tags, 1)
		g.Statsd.Gauge("codagotchi.world.life", float64(g.World.Tick), append(metrics.Tags, "world:"+g.World.Name), 1)
	}
	if g.World.Tick%600 == 0 {
		g.Save(g.SaveName)
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.World.Draw(screen)
	g.Bob.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1080, 720
}

func (g *Game) Save(filename string) {
	data, err := json.Marshal(g)
	if err != nil {
		fmt.Println("error:", err)
	}
	file, err := os.Create(filename)

	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(string(data))
}

func Load(filename string) (*Game, error) {
	var g Game
	data := []byte{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	err = json.Unmarshal(data, &g)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}
	g.SaveName = filename

	return &g, err
}

func Create(filename string) (*Game, error) {
	game := &Game{}

	world := world.World{
		Height: 720,
		Width:  1080,
		Speed:  1,
		Name:   "New Horizon Land",
		Seed:   1337,
		Tick:   0,
	}

	bob := creature.Creature{
		Name:            "Bob",
		Life:            1000,
		LifeSpanCounter: 0,
		PosX:            float64(world.Width / 2),
		PosY:            float64(world.Height / 2),
		VelocityX:       0,
		VelocityY:       0,
	}

	game.World = &world
	game.Bob = &bob
	game.SaveName = filename

	return game, nil
}

func Start(game *Game) {
	statsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Println(err)
		log.Println("Continuing anyway...")
	}
	game.Statsd = statsd

	// Assets loading done in the Start function and not in the Create or Load because
	// it's not saved in the save.json file.
	petPng, _, err := image.Decode(bytes.NewReader(assets.PetV1_png))
	if err != nil {
		log.Fatal(err)
	}
	bobImg := ebiten.NewImageFromImage(petPng)
	game.Bob.Image = bobImg

	ebiten.SetWindowSize(game.World.Width, game.World.Height)
	ebiten.SetWindowTitle(game.World.Name)

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
