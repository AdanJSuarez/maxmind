package report

import (
	"github.com/AdanJSuarez/maxmind/internal/logparser"
	"github.com/oschwald/geoip2-golang"
)

//go:generate mockery --inpackage --name=logParser
//go:generate mockery --inpackage --name=geoInfo

type logParser interface {
	Parse(line string) (logparser.Log, error)
}

type geoInfo interface {
	OpenDB() error
	GetIPInfo(IPString string) *geoip2.City
}
