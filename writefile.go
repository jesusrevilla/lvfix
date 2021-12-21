package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

var path = "result.csv"

//var data = [][]string{{"Time Stamp", "Event"}}

func writeFile(event string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var data [][]string
	timeStamp := time.Now()
	newData := event
	layout := "2006-01-02 15:04:05"
	data = append(data, []string{timeStamp.Format(layout), newData})

	writer := csv.NewWriter(file)
	writer.WriteAll(data)

	if err != nil {
		log.Fatal(err)
	}
}
