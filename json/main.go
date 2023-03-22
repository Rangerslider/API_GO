package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platfrom string
	Password string
	Tags     []string
}

func main() {
	//jsonencode()
	decodejson()
}

func decodejson() {
	jsonDataFromWeb := []byte(` 
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "joydhiakry.in",
		"tags": ["web-dev","js"]
	}
	`)
	var joy course
	check := json.Valid(jsonDataFromWeb)
	if check {
		json.Unmarshal(jsonDataFromWeb, &joy)
		fmt.Printf("%#v", joy)
	} else {
		fmt.Println("error")
	}

	var joy1 map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &joy1)
	for key, value := range joy1 {
		fmt.Println(key, "=", value)
	}
}

func jsonencode() {

	joy := []course{
		{"ReactJS Bootcamp", 199, "joydhiakry.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "joydhiakry.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 399, "joydhiakry.in", "hit123", nil},
	}
	response, err := json.MarshalIndent(joy, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", response)
}
