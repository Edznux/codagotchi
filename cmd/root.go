package cmd

import (
	"fmt"
	"os"

	"github.com/edznux/codagotchi/game"
	"github.com/spf13/cobra"
)

var saveFile string
var rootCmd = &cobra.Command{
	Use:   "codagotchi",
	Short: "A mini coding based game",
	Run: func(cmd *cobra.Command, args []string) {
		var g *game.Game
		var err error

		if _, err = os.Stat(saveFile); err == nil {
			g, err = game.Load(saveFile)
			if err != nil {
				fmt.Println("Error in loading save file", err)
				return
			}
		} else {
			g, err = game.Create(saveFile)
			if err != nil {
				fmt.Println("Error in loading save file", err)
				return
			}
		}

		game.Start(g)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&saveFile, "save", "save.json", "Filename of the save. Default is save.json")
}
