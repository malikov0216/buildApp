package cmd

import (
	"strings"
)
func GetDirectorySrc (src string) string{
	directories := strings.Split(src, "/")
	directoriesArray := directories[0:len(directories) - 1]
	joinedDirectories := strings.Join(directoriesArray, "/")
	directory := "./" + joinedDirectories
	return directory
}