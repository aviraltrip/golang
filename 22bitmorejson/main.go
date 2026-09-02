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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev", "js"]
	}
`)
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
	// some cases where just wanna add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}

}