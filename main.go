package main

import (
	"fmt"
	getChanged "github.com/malikov0216/get-lerna-changed"
	hasJSON "github.com/malikov0216/if-json-has"
	"io/ioutil"
	"log"
	"os/exec"
	"path"
	"strings"
)

func main () {
	projectName, err := getChanged.ExecuteLerna()
	if err != nil {
		log.Fatalln(err)
	}
	buildApp("./packages", projectName)

}
func buildApp(src, projectName string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Println(nil)
	}

	for _, f := range files {
		newSrc := path.Join(src, f.Name())

		if f.IsDir() && f.Name() != "node_modules" {
			buildApp(newSrc, projectName)
		} else if !f.IsDir() && f.Name() != "package.json" || f.IsDir() && f.Name() == "node_modules" {
			continue
		} else if f.Name() == "package.json" {

			if hasDep := hasJSON.FileRead(newSrc, projectName); hasDep {
				dirArr := strings.Split(newSrc, "/")
				dir := dirArr[0:len(dirArr)-1]
				str := "./" + strings.Join(dir, "/")
				cmd := exec.Command("yarn", "build")
				cmd.Dir = str
				out, err := cmd.Output()
				if err != nil {
					log.Fatalln(err)
				}
				a := string(out)
				b := strings.Split(a, "\n")
				for _, k := range b {
					fmt.Println(k)
				}
			}
		}
	}

}