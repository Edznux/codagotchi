package game

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/edznux/codagotchi/game/creature"
	"github.com/edznux/codagotchi/game/world"
	"github.com/edznux/codagotchi/metrics"
)

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
	// Keep an eye on the save size. Should not grow too much
	metrics.Gauge("codagotchi.save.size", float64(len(data)), metrics.Tags, 1)

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
