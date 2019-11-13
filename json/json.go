package main

import (
    "encoding/json"
	"fmt"
)

type superhero struct {
	Name string
	HeroName string
    Powers []string
}


func main() {
	
	batman :=  &superhero{
		Name: "Bruce Wayne",
		HeroName: "Batman",
		Powers: []string{"money", "preparing"}}

	batJson, _ := json.Marshal(batman)
	fmt.Println(string(batJson))

	superman := `{"name": "Clark Kent or Kal-El", "heroname": "Superman", "powers": ["super strength", "heat vision", "x-ray vision", "fly", "super speed", "american way of life"]}`
	supermanStruct := superhero{}
    json.Unmarshal([]byte(superman), &supermanStruct)
    fmt.Println(supermanStruct)
}