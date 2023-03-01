package geoinfo

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
)

const defaultDBPath = "./GeoLite2-City.mmdb"

type GeoInfo struct {
	db *geoip2.Reader
}

func New() *GeoInfo {
	return &GeoInfo{}
}

func (gi *GeoInfo) OpenDB(path string) error {
	if len(path) == 0 {
		path = defaultDBPath
	}

	db, err := geoip2.Open(path)
	if err != nil {
		return fmt.Errorf("error on opening the db: %s: %v", path, err)
	}

	gi.db = db
	return nil
}

func (gi *GeoInfo) GetIPInfo(IPString string) *geoip2.City {
	IP := net.ParseIP(IPString)
	record, err := gi.db.City(IP)
	if err != nil {
		return nil
	}
	return record
}
