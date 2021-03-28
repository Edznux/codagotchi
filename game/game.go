package game

import (
	"encoding/json"
	_ "image/png"
	"log"
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

	// Gather metrics of the game every 10 sec
	if g.World.Tick%600 == 0 {
		metrics.Gauge("codagotchi.bob.lifespan", float64(g.Bob.LifeSpanCounter), metrics.Tags, 1)
		metrics.Gauge("codagotchi.bob.life", float64(g.Bob.Life), metrics.Tags, 1)
		metrics.Gauge("codagotchi.world.tick", float64(g.World.Tick), append(metrics.Tags, "world:"+g.World.Name), 1)
		metrics.Gauge("codagotchi.world.tps", ebiten.CurrentTPS(), append(metrics.Tags, "world:"+g.World.Name), 1)
	}

	if g.World.Tick%60 == 0 {
		g.Save(g.SaveName)
	}

	if g.World.Tick%100 == 0 {
		g.Bob.TargetX, g.Bob.TargetY = g.World.GetRandomPosInBound()
	}

	// Update our main Character and the world at the end
	g.Bob.Update(g.World.Tick)
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
