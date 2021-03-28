// +build wasm

package main

import (
	"fmt"
	"github.com/edznux/codagotchi/game"
)

func main() {
	// This allows to fetch the save over http in the browser when compiled to WASM
	g, err := game.LoadRemote("/save.json")
	if err != nil {
		fmt.Println("Couldn't load save ", err)
	}
	g.Start()
}
