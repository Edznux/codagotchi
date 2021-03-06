package cmd

import (
	"log"

	"github.com/edznux/codagotchi/game"
	"github.com/edznux/codagotchi/webserver"
	"github.com/spf13/cobra"
)

var saveFile string
var rootCmd = &cobra.Command{
	Use:   "codagotchi",
	Short: "A mini coding based game",
	Run: func(cmd *cobra.Command, args []string) {
		g, err := game.LoadOrCreate(saveFile)
		if err != nil {
			log.Printf("Error while creating game: %+v\n", err)
		}
		g.Start()
	},
}
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "A mini coding based game",
	Run: func(cmd *cobra.Command, args []string) {
		web := webserver.WebServer{}
		web.Start(saveFile)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(webCmd)
	rootCmd.PersistentFlags().StringVar(&saveFile, "save", "save.json", "Filename of the save. Default is save.json")
}
