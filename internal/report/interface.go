package report

import (
	geoinfo "github.com/AdanJSuarez/maxmind/internal/geoinfo"
	"github.com/AdanJSuarez/maxmind/internal/logparser"
)

//go:generate mockery --inpackage --name=logParser
//go:generate mockery --inpackage --name=geoInfo

type logParser interface {
	Parse(line string) (logparser.Log, error)
}

type geoInfo interface {
	OpenDB() error
	GetIPInfo(IPString string) geoinfo.GeoInfoModel
}
