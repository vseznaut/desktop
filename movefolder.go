package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func moveFolder(folderName string) {

	// get the current directory
	folderName = getHomeDir() + "/" + folderName

	// read all files in the folder
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}

	// loop all files
	for _, file := range files {

		fi, err := os.Stat(folderName + "/" + file.Name())
		if err != nil {
			return
		}

		// understood modification time
		modTime := fi.ModTime()
		fmt.Println("Last modified time : ", modTime)

		// understood if file is hidden or system
		isHidden, _ := isHidden(file.Name())
		fmt.Println(file.Name(), file.IsDir(), isHidden)

		newPath := ""

		// if file is NOT hidden or system, move to the DROPBOX folder
		if isHidden == false {
			newPath = makeDropboxDir(modTime)
			moveFiles(folderName, newPath, file.Name())
		}
	}

}
