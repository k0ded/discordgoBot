package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func LoadQuotes(fileName string) *[]string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("error loading config,", err)
		return nil
	}
	var quotes []string

	err = json.Unmarshal(content, &quotes)
	if err != nil {
		log.Fatal(fmt.Sprintf("error loading json in %v,", fileName), err)
		return nil
	}

	return &quotes
}
