package geoinfo

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

//go:generate mockery --inpackage --name=geoIPReader

type geoIPDBReader interface {
	City(IP net.IP) (*geoip2.City, error)
	Close() error
}
