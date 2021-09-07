package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string
	Age uint
}

func newPerson(name string, age uint) *Person {
	return &Person{
		Name: name,
		Age: age,
	}
}

func main() {
	filename := "a.txt"
	f, err := os.Create(filename)
	if err != nil {

	}
	defer f.Close()

	decodeJSON := json.NewEncoder(f)
	err = decodeJSON.Encode(newPerson("bruce", 18))
	if err != nil {

	}

}
