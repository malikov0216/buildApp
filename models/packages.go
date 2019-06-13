package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Packages struct {
	Dependencies map[string]string
}

func (packageJSON Packages)IsImported (src, projectName string) bool {
	jsonFile, err := os.Open(src)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &packageJSON)
	if err != nil {
		log.Fatalln(err)
	}
	for key := range packageJSON.Dependencies {
		if key == projectName {
			return true
		} else {
			return false
		}
	}
	return false
}