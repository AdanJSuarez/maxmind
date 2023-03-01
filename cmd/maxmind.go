package main

import (
	"log"
	"time"

	"github.com/AdanJSuarez/maxmind/internal"
)

const channelSize = 10

func main() {
	log.Println("Start")
	linesCh := make(chan string, channelSize)
	reader := internal.New(linesCh)
	go reader.ReadLinesFromFile("./asset/access.log")
	for line := range linesCh {
		log.Println(line)
		time.Sleep(3 * time.Second)
	}
}
