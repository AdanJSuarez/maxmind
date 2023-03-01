package geoinfo

import "github.com/oschwald/geoip2-golang"

type GeoInfo struct {
	db string
}

func New() *GeoInfo {
	db, err := geoip2.Open("GeoIP2-City.mmdb")
	return &GeoInfo{}
}
