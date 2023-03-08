package report

import (
	"fmt"
	"regexp"

	"github.com/AdanJSuarez/maxmind/internal/report/countries"
)

const (
	logPatter         = `[a-f0-9]+/css/|\w\.css/?$|[a-f0-9]+/images/|/images/|[a-f0-9]+/js/|\w\.js/?$|/entry-images/|/static/|/robots.txt/?$|/favicon.ico/?$|\w\.rss/?$|\w\.atom/?$`
	excludedPage      = "/"
	unitedStates      = "United States"
	top               = 10
	countriesDataName = "Countries"
)

type Report struct {
	regex *regexp.Regexp
	data  countriesData
}

// New returns an instance of Report initialized.
func New() *Report {
	report := &Report{}
	report.regex = regexp.MustCompile(logPatter)
	report.data = countries.New(countriesDataName)

	return report
}

// Generate prints reports of Countries and US.
func (r *Report) Generate() {
	fmt.Println("\n==> World and US report:")
	r.printReport()
}

// ShouldExclude returns true is the page should be excluded from the report.
func (r *Report) ShouldExclude(page string) bool {
	return r.regex.MatchString(page)
}

// Subdivision returns the first element of subdivisions.
func (r *Report) Subdivision(subdivisions []string) string {
	return subdivisions[0]
}

// AddData adds a new element to data.
func (r *Report) AddData(countryName, subdivisionName, pageName string) {
	r.data.AddToCountries(countryName, subdivisionName, pageName)
}

func (r *Report) printReport() {
	r.printCountries()
	r.printUSA()
}

func (r *Report) printCountries() {
	fmt.Println("==> Countries:")
	countries := r.data.TopAreas(r.data.Name(), excludedPage, top)
	for idx, country := range countries {
		fmt.Printf("%d : %s - Visits: %d - Most visited page: \"%s\"\n", idx+1, country.Name(), country.Visit(), country.TopPage())
	}
}

func (r *Report) printUSA() {
	fmt.Println("==> United States:")
	us := r.data.TopAreas(unitedStates, excludedPage, top)
	for idx, state := range us {
		fmt.Printf("%d: %s - Visits: %d - Most visited page: \"%s\"\n", idx+1, state.Name(), state.Visit(), state.TopPage())
	}
}
