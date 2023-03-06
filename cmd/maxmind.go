package main

import (
	"fmt"
	"log"
	"sync"

	geoinfo "github.com/AdanJSuarez/maxmind/internal/geoinfo"
	"github.com/AdanJSuarez/maxmind/internal/logparser"
	"github.com/AdanJSuarez/maxmind/internal/logreader"
	"github.com/AdanJSuarez/maxmind/internal/report"
)

const (
	channelSize = 1000
	fileName    = "./asset/GeoLite2-City.mmdb"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("==> Start <==")
	logParser, err := logparser.New()
	if err != nil {
		log.Panicf("error on log parser: %v", err)
	}

	geoinfo := geoinfo.New(fileName)
	if err := geoinfo.OpenDB(); err != nil {
		log.Panicf("error open db: %v", err)
	}
	defer geoinfo.Close()

	report := report.New(logParser, geoinfo, channelSize, &wg)

	reader := logreader.New(report.LinesCh())

	fmt.Printf("==> Reading records from file: %s\n", fileName)
	go reader.ReadLinesFromFile("./asset/access.log")
	wg.Add(1)
	go report.GetReport()
	wg.Wait()

}
