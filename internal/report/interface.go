package report

import "github.com/AdanJSuarez/maxmind/internal/report/countries"

//go:generate mockery --inpackage --name=countriesData

type countriesData interface {
	Name() string
	AddToCountries(countryName, subdivisionName, webpageName string)
	TopAreas(name, pageExcluded string, topNumber int) []countries.Info
}
