package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Story map[string]PlotPoint

type PlotPoint struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []PlotOptions `json:"options"`
}

type PlotOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	story := getJsonFileAsMap("gopher.json")

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("plot-point.html"))
		tmpl.Execute(w, story["intro"])
	})

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", nil)
}

func getJsonFileAsMap(filename string) Story {
	content, err := ioutil.ReadFile(filename)

	var story Story
	err = json.Unmarshal(content, &story)
	if err != nil {
		panic(err)
	}

	return story
}
