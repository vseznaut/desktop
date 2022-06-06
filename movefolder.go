package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func moveFolder(folderName string) {

	// get downloads dir
	folderName = getHomeDir() + "/" + folderName

	// read downloads dir
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}

	// print filename
	for _, file := range files {

		fi, err := os.Stat(folderName + "/" + file.Name())
		if err != nil {
			return
		}
		modTime := fi.ModTime()
		fmt.Println("Last modified time : ", modTime)

		isHidden, _ := isHidden(file.Name())
		fmt.Println(file.Name(), file.IsDir(), isHidden)

		newPath := ""

		if isHidden == false {
			newPath = makeDropboxDir(modTime)
			moveFiles(folderName, newPath, file.Name())
		}
	}

}
