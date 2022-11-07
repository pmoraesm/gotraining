package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var inputName string
	var inputAddress string

	fmt.Printf("Please input your name:\n")
	fmt.Scan(&inputName)
	fmt.Printf("Please input your address:\n")
	fmt.Scan(&inputAddress)

	Person := map[string]string{"name": inputName,
		"address": inputAddress}

	jPerson, err := json.Marshal(Person)
	if err == nil {
		fmt.Println(string(jPerson))
	}
}
