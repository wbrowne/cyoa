package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	fmt.Print(story)
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
