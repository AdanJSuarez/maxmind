package report

import (
	"fmt"
	"log"
	"regexp"
	"sync"

	"github.com/AdanJSuarez/maxmind/internal/report/countries"
)

const (
	cssOr         = `[a-f0-9]+/css/|`
	imagesOr      = `[a-f0-9]+/images/|/images/|`
	entryImagesOr = `/entry-images/|`
	staticOr      = `/static/|`
	robotstxtOr   = `/robots.txt/?$|`
	faviconicoOr  = `/favicon.ico/?$|`
	rssOr         = `\w\.rss/?$|`
	atomOr        = `\w\.atom/?$`
	logPatter     = cssOr + imagesOr + entryImagesOr + staticOr + robotstxtOr + faviconicoOr + rssOr + atomOr
)

type Report struct {
	wg        *sync.WaitGroup
	logParser logParser
	geoInfo   geoInfo
	regex     *regexp.Regexp
	countries *countries.Countries
	linesCh   chan string
}

// New returns an instance of Report initialized with logParser, geoInfo and the linesCh size.
func New(logParser logParser, geoInfo geoInfo, channelSize int64, wg *sync.WaitGroup) *Report {
	return &Report{
		wg:        wg,
		regex:     regexp.MustCompile(logPatter),
		logParser: logParser,
		geoInfo:   geoInfo,
		countries: countries.NewCountries(),
		linesCh:   make(chan string, channelSize),
	}
}

// LinesCh returns the channel for log lines
func (r *Report) LinesCh() chan string {
	return r.linesCh
}

func (r *Report) GetReport() {
	defer r.wg.Done()
	for line := range r.linesCh {
		lineLog, err := r.logParser.Parse(line)
		if err != nil {
			log.Printf("log line excluded: %v\n", err)
		}

		if r.shouldExclude(lineLog.RequestPath) {
			// log.Printf("log excluded in report: %v", lineLog)
			continue
		}
		record := r.geoInfo.GetIPInfo(lineLog.IP)
		// log.Println("Info ==> ", record)
		if len(record.Subdivisions) > 0 {
			r.countries.AddToCountries(record.Country.Names["en"], record.Subdivisions[0].Names["en"], lineLog.RequestPath)
		} else {
			r.countries.AddToCountries(record.Country.Names["en"], "unknown", lineLog.RequestPath)
		}

	}
	for idx, val := range r.countries.MostVisit(10) {
		fmt.Printf("Number: %d: %s\n", idx+1, val.Name())
	}

}

func (r *Report) shouldExclude(requestPath string) bool {
	return r.regex.MatchString(requestPath)
}

func (r *Report) extractSubdivisions(subdivision []interface{}) {
	// TODO: Implement extractSubdivisions
}
