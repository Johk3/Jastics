package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
)


type Feedback struct {
	ratios []float64
	text []string
}

func (f *Feedback) AppendRatios(ratios float64) {
	f.ratios = append(f.ratios, ratios)
}

func (f *Feedback) AppendText(text string) {
	f.text = append(f.text, text)
}

// -- -- --


func Run(input string, output string) {
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
						modelF.AppendText(v.Elem().MapIndex(f).Elem().String())
					}
					if f.String() == "ratios" {
						for z := 0; z < v.Elem().MapIndex(f).Elem().Len(); z++ {
							ratioS := v.Elem().MapIndex(f).Elem().Index(z).Elem().String()
							if ratioS == "<float64 Value>" {
								modelF.AppendRatios(v.Elem().MapIndex(f).Elem().Index(z).Elem().Float())
							} else {
								if s, err := strconv.ParseFloat(ratioS, 64); err == nil {
									modelF.AppendRatios(s)
								}
							}
						}
					}
				}
			}
		}
		fmt.Printf("%d/%d\n", i, _result.Len())
	}

	// Output feedback, for python sentiment analysis
	writeSentiments(&modelF)

	// Analyze the statistical data of ratios, and graph them
	analyzeRatios(&modelF)
}

func analyzeRatios(m *Feedback) {
	var meanAverage int

	for n := range m.ratios {
		meanAverage += n
	}
	meanAverage = meanAverage/len(m.ratios)
	// Convert to percentages
	meanAverage = meanAverage/5

	fmt.Printf("User experience ranges from 100-500\nThe average user experience was: %d%%\n", meanAverage)
}

func writeSentiments(m *Feedback) {
	fmt.Println("Writing down sentiments for python...")

	feedbackPath := "/tmp/jm_feedback"
	var bytes int
	f, err := os.Create(feedbackPath)
	check(err)

	defer f.Close()

	for _, text := range m.text {
		n3, err := f.WriteString(text + "-+--++--+-+--++--+--+-+--+")
		check(err)
		bytes += n3
	}
	f.Sync()

	fmt.Printf("Wrote %d bytes\n", bytes)
}

func check(err error) {
	if err != nil {
		log.Fatalln("Jastics error: ", err)
	}
}
