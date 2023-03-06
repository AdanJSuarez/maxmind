package report

import (
	"fmt"
	"log"
	"regexp"
	"sync"

	"github.com/AdanJSuarez/maxmind/internal/report/countries"
)

const (
	logPatter    = `[a-f0-9]+/css/|[a-f0-9]+/images/|/images/|[a-f0-9]+/js/|/entry-images/|/static/|/robots.txt/?$|/favicon.ico/?$|\w\.rss/?$|\w\.atom/?$`
	excludedPage = "/"
	unitedStates = "United States"
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
			continue
		}
		record := r.geoInfo.GetIPInfo(lineLog.IP)
		r.countries.AddToCountries(record.CountryName, record.Subdivisions[0], lineLog.RequestPath)
	}
	r.printReport()
}

func (r *Report) shouldExclude(requestPath string) bool {
	return r.regex.MatchString(requestPath)
}

func (r *Report) printReport() {
	fmt.Println("==> Countries:")
	for idx, val := range r.countries.TopAreas(r.countries.Countries(), excludedPage) {
		fmt.Printf("%d : %s - Visits: %d - Most visited page: \"%s\"\n", idx+1, val.Name, val.Visit, val.TopPage)
	}
	fmt.Println("==> United States:")
	for idx, val := range r.countries.TopAreas(r.countries.Countries().Children()[unitedStates], excludedPage) {
		fmt.Printf("%d: %s - Visits: %d - Most visited page: \"%s\"\n", idx+1, val.Name, val.Visit, val.TopPage)
	}
	fmt.Println("==> Finished <== ")
}
