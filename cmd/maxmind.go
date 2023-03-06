package main

import (
	"fmt"
	"sync"

	"github.com/AdanJSuarez/maxmind/internal/configuration"
	geoinfo "github.com/AdanJSuarez/maxmind/internal/geoinfo"
	"github.com/AdanJSuarez/maxmind/internal/logparser"
	"github.com/AdanJSuarez/maxmind/internal/logreader"
	"github.com/AdanJSuarez/maxmind/internal/report"
)

const (
	channelSize = 1000
)

func main() {
	config := configuration.New()
	var wg sync.WaitGroup

	logParser, err := logparser.New()
	if err != nil {
		fmt.Printf("error on log parser: %v\n", err)
		return
	}

	geoInfo := geoinfo.New(config.DBfile)
	if err := geoInfo.OpenDB(); err != nil {
		fmt.Printf("error open db: %v\n", err)
		return
	}
	defer geoInfo.Close()

	report := report.New(logParser, geoInfo, channelSize, &wg)

	reader := logreader.New(report.LinesCh())

	fmt.Printf("==> Reading logs from file: %s\n", config.LogFile)
	go reader.ReadLinesFromFile(config.LogFile)
	wg.Add(1)
	go report.GetReport()
	wg.Wait()
}
