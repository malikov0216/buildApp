package cmd

import (
	"github.com/malikov0216/buildApp/models"
	"io/ioutil"
	"path"
)

//building App
func BuildApp(src, projectName string) error{
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {
		newSrc := path.Join(src, f.Name())

		//Conditions
		noPackage := !f.IsDir() && f.Name() != "package.json" || f.IsDir() && f.Name() == "node_modules"
		isPackage := f.Name() == "package.json"
		isDirectory := f.IsDir() && f.Name() != "node_modules"

		var packages models.Packages
		//Check Conditions
		if isDirectory {
			if err := BuildApp(newSrc, projectName); err != nil {
				return err
			}
		} else if noPackage {
			continue
		} else if isPackage {
			packages = packages.Read(newSrc, projectName)
			if packages.IsImported {
				srcToBuild := GetDirectorySrc(newSrc)
				if err := ExecuteBuild(srcToBuild); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
