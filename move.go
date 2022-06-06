package main

import (
	"fmt"
	"os"
)

func moveFiles(oldPath string, newPath string, fileName string) {
	fmt.Println("oldPath: " + oldPath)
	fmt.Println("newPath: " + newPath)
	fmt.Println("fileName: " + fileName)

	err := os.Rename(oldPath+"/"+fileName, newPath+"/"+fileName) // rename directory

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}
