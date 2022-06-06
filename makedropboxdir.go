package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func makeDropboxDir(currentTime time.Time) (workDay string) {

	dropbox := getHomeDir() + "/Dropbox"
	workMonth := dropbox + "/WORK-" + currentTime.Format("2006-01")
	fmt.Println(workMonth)

	_, err := os.Stat(workMonth)
	if err != nil {
		// File or directory does not exist
		if os.IsNotExist(err) {
			if err := os.Mkdir(workMonth, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}
	}

	workDay = workMonth + "/" + currentTime.Format("2006-01-02")
	fmt.Println(workDay)

	_, err = os.Stat(workDay)
	if err != nil {
		// File or directory does not exist
		if os.IsNotExist(err) {
			if err := os.Mkdir(workDay, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}
	}

	return workDay
}
