package report

import (
	"fmt"
	"regexp"

	"github.com/AdanJSuarez/maxmind/internal/report/countries"
)

const (
	logPatter    = `[a-f0-9]+/css/|[a-f0-9]+/images/|/images/|[a-f0-9]+/js/|/entry-images/|/static/|/robots.txt/?$|/favicon.ico/?$|\w\.rss/?$|\w\.atom/?$`
	excludedPage = "/"
	unitedStates = "United States"
	top          = 10
)

type Report struct {
	regex *regexp.Regexp
	data  countriesData
}

// New returns an instance of Report initialized with logParser, geoInfo and the linesCh size.
func New() *Report {
	report := &Report{}
	report.regex = regexp.MustCompile(logPatter)
	report.data = countries.New()

	return report
}

func (r *Report) Generate() {
	r.printReport()
}

func (r *Report) ShouldExclude(requestPath string) bool {
	return r.regex.MatchString(requestPath)
}

func (r *Report) Subdivision(subdivisions []string) string {
	return subdivisions[0]
}

func (r *Report) AddData(countryName, subdivisionName, pageName string) {
	r.data.AddToCountries(countryName, subdivisionName, pageName)
}

func (r *Report) printReport() {
	r.printCountries()
	r.printUSA()
}

func (r *Report) printCountries() {
	fmt.Println("==> Countries:")
	for idx, val := range r.data.TopAreas(r.data.Name(), excludedPage, top) {
		fmt.Printf("%d : %s - Visits: %d - Most visited page: \"%s\"\n", idx+1, val.Name, val.Visit, val.TopPage)
	}
}

func (r *Report) printUSA() {
	fmt.Println("==> United States:")
	for idx, val := range r.data.TopAreas(unitedStates, excludedPage, top) {
		fmt.Printf("%d: %s - Visits: %d - Most visited page: \"%s\"\n", idx+1, val.Name, val.Visit, val.TopPage)
	}
}
