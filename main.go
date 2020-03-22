package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
)

const (
    STATIC_DIR = "/static/"
    PORT       = "8080"
)

type Story map[string]PlotPoint

type PlotPoint struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []PlotOptions `json:"options"`
}

type PlotOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	story := importJsonAsStory("gopher.json")

	tmpl := template.Must(template.ParseFiles("layout.html", "plot-point.html"))

	router := mux.NewRouter()

    // Static assets
    router.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	router.HandleFunc("/", func(w http.ResponseWriter, router *http.Request) {
		http.Redirect(w, router, "/intro", http.StatusFound)
	})

	router.HandleFunc("/{arc}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		arc := vars["arc"]

		tmpl.Execute(w, story[arc])
	})

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", router)
}

func importJsonAsStory(filename string) Story {
	content, err := ioutil.ReadFile(filename)

	var story Story
	err = json.Unmarshal(content, &story)
	if err != nil {
		panic(err)
	}

	return story
}
