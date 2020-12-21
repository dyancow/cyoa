package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//can make a map of story_arcs
type stories map[string]Story_arc

//the fields in the structs for json must be upper case to be properly exported.
type Story_arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func load_story_data() (loaded_stories stories) {
	//load JSON file into memory. Check for load error.
	jsonFile, err := os.Open("./gopher.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	//convert JSON file data to byte array.
	jsonByte, _ := ioutil.ReadAll(jsonFile)

	//make a variable of structs to hold the information from the json byte array.
	var stories stories
	jsonerr := json.Unmarshal(jsonByte, &stories)
	if jsonerr != nil {
		fmt.Println(err)
		return
	}
	return stories
}

func main() {
	loaded_stories := load_story_data()

	//output check
	fmt.Println("Checking how we unbundled the json file: ")

	for key, arc := range loaded_stories {
		fmt.Printf("Current story is: %s with the title of %s\n", key, arc.Title)
	}
}
