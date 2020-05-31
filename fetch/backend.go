package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"reflect"
)


func Run(input string, output string) map[int]int {
	results := make(map[int]int)

	fmt.Printf("Statistically analyzing %v, and dumping to %v\n", input, output)
	jsonFile, err := os.Open(input)
	check(err)

	// defer close the json file so it can be parsed later
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	check(err)

	_result := reflect.ValueOf(result)
	for i := 0; i < _result.Len(); i++ {
		for _, e := range _result.MapKeys() {
			v := _result.MapIndex(e)
			fmt.Println(v)
		}
		break
	}

	return results
}

func check(err error) {
	if err != nil {
		log.Fatalln("Jastics error: ", err)
	}
}
