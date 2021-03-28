package game

import (
	"encoding/json"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/edznux/codagotchi/game/creature"
	"github.com/edznux/codagotchi/game/gui"
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

	GUI gui.GUI `json:"-"`

	// Don't save the Statsd client in the JSON of the save
	Statsd *statsd.Client `json:"-"`
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	if g.World.Tick%600 == 0 {
		metrics.Gauge("codagotchi.bob.lifespan", float64(g.Bob.LifeSpanCounter), metrics.Tags, 1)
		metrics.Gauge("codagotchi.bob.life", float64(g.Bob.Life), metrics.Tags, 1)
		metrics.Gauge("codagotchi.world.tick", float64(g.World.Tick), append(metrics.Tags, "world:"+g.World.Name), 1)
	}
	if g.World.Tick%60 == 0 {
		g.Save(g.SaveName)
	}

	if g.World.Tick%100 == 0 {
		g.Bob.TargetX, g.Bob.TargetY = g.World.GetRandomPosInBound()
	}

	g.Bob.Update()
	g.World.Update()

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.World.Draw(screen)
	g.Bob.Draw(screen)
	g.GUI.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1080, 720
}

func (g *Game) Save(filename string) {
	data, err := json.Marshal(g)
	if err != nil {
		log.Println("error:", err)
	}
	file, err := os.Create(filename)

	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(string(data))
}

func LoadRemote(url string) (*Game, error) {
	var g Game
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Unable to get remote file: %v", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(data, &g)
	if err != nil {
		log.Println("JSON Unmarshal error:", err)
		return nil, err
	}
	g.SaveName = "save.json"

	return &g, err
}

func Load(filename string) (*Game, error) {
	var g Game
	data := []byte{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Unable to read file: %v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &g)
	if err != nil {
		log.Println("JSON Unmarshal error:", err)
		return nil, err
	}
	g.SaveName = filename

	return &g, err
}

func Create(filename string) (*Game, error) {
	game := Game{}

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
		Speed:           1,
		VelocityX:       0,
		VelocityY:       0,
	}

	game.World = &world
	game.Bob = &bob
	game.SaveName = filename
	game.Save(game.SaveName)
	return &game, nil
}

func (game *Game) Start() {

	game.World.Init()
	game.Bob.Init()

	ebiten.SetWindowSize(game.World.Width, game.World.Height)
	ebiten.SetWindowTitle(game.World.Name)

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func LoadOrCreate(saveFile string) (*Game, error) {
	var g *Game

	_, err := os.Stat(saveFile)
	if err == nil {
		log.Println("Trying to load file", saveFile)
		g, err = Load(saveFile)
		if err != nil {
			log.Println("Error in loading save file", err)
			return nil, err
		}
	} else {
		log.Println("Trying to create file", saveFile)
		g, err = Create(saveFile)
		if err != nil {
			log.Println("Error in creating save file", err)
			return nil, err
		}
	}
	return g, nil
}
