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
	jsOr          = `[a-f0-9]+/js/|`
	entryImagesOr = `/entry-images/|`
	staticOr      = `/static/|`
	robotstxtOr   = `/robots.txt/?$|`
	faviconicoOr  = `/favicon.ico/?$|`
	rssOr         = `\w\.rss/?$|`
	atomOr        = `\w\.atom/?$`
	logPatter     = cssOr + imagesOr + jsOr + entryImagesOr + staticOr + robotstxtOr + faviconicoOr + rssOr + atomOr
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
		countries: countries.New(),
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
		r.countries.AddToCountries(record.CountryName, record.Subdivisions[0], lineLog.RequestPath)

	}
	r.printReport()
}

func (r *Report) shouldExclude(requestPath string) bool {
	return r.regex.MatchString(requestPath)
}

func (r *Report) printReport() {
	fmt.Println("Countries:")
	for idx, val := range r.countries.TopAreas(r.countries.Countries(), "/") {
		fmt.Printf("%d : %s - Visits: %d - Most visited page: %s\n", idx, val.Name, val.Visit, val.TopPage)
	}
	fmt.Println("United States:")
	for _, val := range r.countries.TopAreas(r.countries.Countries().Children()["United States"], "/") {
		fmt.Printf("State: %s - Visits: %d - Most visited page: %s\n", val.Name, val.Visit, val.TopPage)
	}
	log.Println("==> Finished <== ")
}
