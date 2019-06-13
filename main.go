package main

import (
	"github.com/malikov0216/buildApp/cmd"
	"log"
)

func main () {
	projectName, err := cmd.ExecuteLerna()
	if err != nil {
		log.Fatalln(err)
	}
	if err := cmd.BuildApp("./packages", projectName); err != nil {
		log.Fatalln(err)
	}
}


