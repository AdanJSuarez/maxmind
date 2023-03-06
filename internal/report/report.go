package report

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/AdanJSuarez/maxmind/internal/geoinfo"
	"github.com/AdanJSuarez/maxmind/internal/logparser"
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
func New(geoInfo geoInfo, channelSize int64, wg *sync.WaitGroup) *Report {
	report := &Report{}
	report.wg = wg
	report.regex = regexp.MustCompile(logPatter)
	report.logParser = report.setLogParser()
	report.geoInfo = geoInfo
	report.countries = countries.New()
	report.linesCh = make(chan string, channelSize)
	return report
}

// LinesCh returns the channel for log lines
func (r *Report) LinesCh() chan string {
	return r.linesCh
}

func (r *Report) GenerateReport() {
	defer r.wg.Done()

	r.extractDataFromLogs()
	r.printReport()
}

func (r *Report) setLogParser() logParser {
	logParser, err := logparser.New()
	if err != nil {
		fmt.Printf("error on log parser: %v\n", err)
	}
	return logParser
}

func (r *Report) extractDataFromLogs() {
	for line := range r.linesCh {
		lineLog, err := r.logParser.Parse(line)
		if err != nil {
			fmt.Printf("log line excluded: %v\n", err)
			continue
		}

		if r.shouldExclude(lineLog.RequestPath) {
			continue
		}
		record := r.geoInfo.GetIPInfo(lineLog.IP)
		r.countries.AddToCountries(record.CountryName, r.mainSubdivision(record), lineLog.RequestPath)
	}
}

func (r *Report) shouldExclude(requestPath string) bool {
	return r.regex.MatchString(requestPath)
}

func (r *Report) printReport() {
	r.printCountries()
	r.printUSA()
	fmt.Println("==> Finished <== ")
}

func (r *Report) printCountries() {
	fmt.Println("==> Countries:")
	for idx, val := range r.countries.TopAreas(r.countries.Countries().Name(), excludedPage) {
		fmt.Printf("%d : %s - Visits: %d - Most visited page: \"%s\"\n", idx+1, val.Name, val.Visit, val.TopPage)
	}
}

func (r *Report) printUSA() {
	fmt.Println("==> United States:")
	for idx, val := range r.countries.TopAreas(unitedStates, excludedPage) {
		fmt.Printf("%d: %s - Visits: %d - Most visited page: \"%s\"\n", idx+1, val.Name, val.Visit, val.TopPage)
	}
}

func (r *Report) mainSubdivision(record geoinfo.GeoInfoModel) string {
	return record.Subdivisions[0]
}
