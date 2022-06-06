package main

import (
	"fmt"
	"log"
	"os"
)

func getHomeDir() (homeDir string) {

	// Get home dir
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirname)

	return dirname
}
