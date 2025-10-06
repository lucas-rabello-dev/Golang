package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `json:"name"`
	Age int `json:age`
}

func main() {

	jsonString := `{"name":"Rodrigo", "age": 30}`
	var person Person

	json.Unmarshal([]byte(jsonString), &person)
	fmt.Printf("Nome: %s, Idade: %d \n", person.Name, person.Age)


	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))
}