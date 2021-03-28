// +build wasm

package main

import (
	"github.com/edznux/codagotchi/game"
	"log"
)

func main() {
	// This allows to fetch the save over http in the browser when compiled to WASM
	g, err := game.LoadRemote("/save.json")
	if err != nil {
		log.Println("Couldn't load save ", err)
	}
	g.Start()
}
