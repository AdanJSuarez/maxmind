package main

import (
	"log"
	"sync"

	geoinfo "github.com/AdanJSuarez/maxmind/internal/geoinfo"
	"github.com/AdanJSuarez/maxmind/internal/logparser"
	"github.com/AdanJSuarez/maxmind/internal/logreader"
	"github.com/AdanJSuarez/maxmind/internal/report"
)

const channelSize = 1000

func main() {
	var wg sync.WaitGroup

	log.Println("==> Start <==")
	logParser, err := logparser.New()
	if err != nil {
		log.Panicf("error on log parser: %v", err)
	}

	geoinfo := geoinfo.New("./asset/GeoLite2-City.mmdb")
	if err := geoinfo.OpenDB(); err != nil {
		log.Panicf("error open db: %v", err)
	}
	defer geoinfo.Close()

	report := report.New(logParser, geoinfo, channelSize, &wg)

	reader := logreader.New(report.LinesCh())

	go reader.ReadLinesFromFile("./asset/access.log")
	// for line := range linesCh {
	// 	log.Println(line)
	// 	time.Sleep(3 * time.Second)
	// }
	wg.Add(1)
	go report.GetReport()
	wg.Wait()

}
