package main

import (
	"fmt"
	"sync"

	"github.com/AdanJSuarez/maxmind/internal/configuration"
	geoinfo "github.com/AdanJSuarez/maxmind/internal/geoinfo"
	"github.com/AdanJSuarez/maxmind/internal/logreader"
	"github.com/AdanJSuarez/maxmind/internal/report"
)

const (
	channelSize      = 1000
	activeGoroutines = 2
)

func main() {
	config := configuration.New()
	var wg sync.WaitGroup

	geoInfo := geoinfo.New(config.DBfile)
	if err := geoInfo.OpenDB(); err != nil {
		fmt.Printf("error open db: %v\n", err)
		return
	}
	defer geoInfo.Close()

	report := report.New(geoInfo, channelSize, &wg)
	logReader, err := logreader.New(config.LogFile, report.LinesCh())
	if err != nil {
		fmt.Printf("error on log file: %v\n", err)
		return
	}

	wg.Add(activeGoroutines)
	fmt.Printf("==> Reading logs from file: %s\n", config.LogFile)
	go readLogFromFile(&wg, logReader)
	go report.GenerateReport()
	wg.Wait()
}

func readLogFromFile(wg *sync.WaitGroup, logReader *logreader.LogReader) {
	defer wg.Done()
	if err := logReader.ReadLinesFromFile(); err != nil {
		fmt.Println(err)
	}
}
