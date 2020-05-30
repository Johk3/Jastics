package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
)


func run(input string, output string) {
	fmt.Printf("Statistically analyzing %v, and dumping to %v\n", input, output)

	jsonFile, err := os.Open(input)
	check(err)

	// defer close the json file so it can be parsed later
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)
}

func check(err error) {
	if err != nil {
		log.Fatalln("Jastics error: ", err)
	}
}
