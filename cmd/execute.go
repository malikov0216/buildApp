package cmd

import (
	"log"
	"os/exec"
	"strings"
)

//get values when executing "lerna changed"
func ExecuteLerna () (string, error)  {
	out, err := exec.Command("lerna", "changed").Output()
	if err != nil {
		return "[Error in executing]", err
	}
	convertedOut := string(out)
	outputs := strings.Split(convertedOut, "\n")
	return outputs[0], nil
}

//Executing command "yarn build"
func ExecuteBuild (srcToBuild string) error{
	command := exec.Command("yarn", "build")
	command.Dir = srcToBuild
	out, err := command.Output()
	if err != nil {
		return err
	}
	convertedOut := string(out)
	separated := strings.Split(convertedOut, "\n")
	for _, k := range separated {
		log.Println(k)
	}
	return nil
}