package geoinfo

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
)

const (
	defaultDBPath = "./GeoLite2-City.mmdb"
)

type GeoInfoRepository struct {
	dbPath string
	db     geoIPDBReader
}

// New returns an instance of GeoInfoRepository
func New(dbPath string) *GeoInfoRepository {
	gi := &GeoInfoRepository{}

	gi.setDBPath(dbPath)

	return gi
}

// OpenDB opens the db from the path set and return nil. It returns an error otherwise.
func (gi *GeoInfoRepository) OpenDB() error {
	db, err := geoip2.Open(gi.dbPath)
	if err != nil {
		return fmt.Errorf("error on opening the db: %s: %v", gi.dbPath, err)
	}

	gi.db = db
	return nil
}

// Close closes the db and return nil. It returns an error otherwise.
func (gi *GeoInfoRepository) Close() error {
	return gi.db.Close()
}

// GetIPInfo returns an instance of geoip2.City with the info of the IP passed if any.
func (gi *GeoInfoRepository) GetIPInfo(IPString string) GeoInfoModel {
	IP := net.ParseIP(IPString)
	record, err := gi.db.City(IP)
	if err != nil || record == nil {
		return newGeoInfoModel(IPString, nil)
	}
	return newGeoInfoModel(IPString, record)
}

func (gi *GeoInfoRepository) setDBPath(dbPath string) {
	if gi.emptyPath(dbPath) {
		gi.dbPath = defaultDBPath
		return
	}
	gi.dbPath = dbPath
}

func (gi *GeoInfoRepository) emptyPath(path string) bool {
	return len(path) == 0
}
