package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	coursesl := []course{
		{"ReactJS", 299, "https://aviraltrip.vercel.app", "abc123", []string{"web-dev", "js"}},
		{"MERN", 199, "https://aviraltrip.vercel.app", "xyz123", []string{"dev-ops", "js"}},
		{"Angular", 399, "https://aviraltrip.vercel.app", "lmao420", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(coursesl, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
