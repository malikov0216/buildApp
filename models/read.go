package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func (packageJSON Packages)Read (src, projectName string) Packages {
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
			packageJSON.IsImported = true
			return packageJSON
		} else {
			packageJSON.IsImported = false
			return packageJSON
		}
	}
	return packageJSON
}