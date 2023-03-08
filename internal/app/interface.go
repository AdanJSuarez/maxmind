package app

import (
	"github.com/AdanJSuarez/maxmind/internal/geoinfo"
	"github.com/AdanJSuarez/maxmind/internal/logparser"
)

//go:generate mockery --inpackage --name=logParser
//go:generate mockery --inpackage --name=geoInfo
//go:generate mockery --inpackage --name=logReader
//go:generate mockery --inpackage --name=countryReport

type logParser interface {
	Parse(line string) (logparser.Log, error)
}

type geoInfo interface {
	OpenDB() error
	GetIPInfo(IPString string) geoinfo.GeoInfoModel
	Close() error
}

type logReader interface {
	Close() error
	Open() error
	ReadLinesFromFile()
}

type countryReport interface {
	AddData(countryName, subdivisionName, pageName string)
	ShouldExclude(page string) bool
	Generate()
	Subdivision(subdivisions []string) string
}
