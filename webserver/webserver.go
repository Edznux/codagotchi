package webserver

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/edznux/codagotchi/game"
	"github.com/edznux/codagotchi/webserver/templates"
)

type WebServer struct {
	saveFile string
	game     *game.Game
}

type variables struct {
	SaveName string
}

func (web *WebServer) Index(w http.ResponseWriter, r *http.Request) {
	template, err := template.New("Index").Parse(templates.IndexTmpl)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
	v := variables{SaveName: web.game.World.Name}

	err = template.Execute(w, v)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

func (web *WebServer) HandleSave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	g, err := game.Load(web.saveFile)
	log.Println("Loading game:", web.saveFile)
	if err != nil {
		log.Println("Error loading game:", err)
		w.Write([]byte{})
	}

	res, err := json.Marshal(g)
	if err != nil {
		log.Println("Error marshalling game save:", err)
		w.Write([]byte{})
	}
	w.Write(res)
}

func (web *WebServer) Start(saveFile string) {
	web.saveFile = saveFile
	g, err := game.LoadOrCreate(web.saveFile)
	if err != nil {
		log.Println("Error loading game:", err)
		return
	}
	web.game = g

	http.HandleFunc("/", web.Index)
	http.HandleFunc("/save.json", web.HandleSave)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./webserver/static/"))))

	go http.ListenAndServe(":8080", nil)

	web.game.Start()
}
