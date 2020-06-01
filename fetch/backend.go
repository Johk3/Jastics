package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"reflect"
)


type Feedback struct {
	ratios []int
	text []string
}

func (f *Feedback) AppendRatios(ratios []int) {
	f.ratios = append(f.ratios, ratios...)
}

func (f *Feedback) AppendText(text string) {
	f.text = append(f.text, text)
}

// -- -- --


func Run(input string, output string) *Feedback {
	modelF := Feedback{}

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
		for _, e := range _result.Index(i).Elem().MapKeys() {
			// Handle operations to fetch ratio and feedbacktext from feedback
			if e.String() == "feedback" {
				v := _result.Index(i).Elem().MapIndex(e)
				for _, f := range v.Elem().MapKeys() {
					if f.String() == "mainFormText" {
						modelF.AppendText(v.Elem().MapIndex(f).String())
					}
					if f.String() == "ratios" {
						fmt.Println(v.Elem().MapIndex(f))
					}
				}
			}
		}
		fmt.Printf("%d/%d\n", i, _result.Len())
		break
	}

	return &modelF
}

func check(err error) {
	if err != nil {
		log.Fatalln("Jastics error: ", err)
	}
}
