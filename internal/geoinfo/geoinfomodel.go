package geoinfo

import "github.com/oschwald/geoip2-golang"

// GeoInfoModel represents the information about a specific IP
type GeoInfoModel struct {
	IP           string
	CountryName  string
	Subdivisions []string
}

func newGeoInfoModel(IPString string, record *geoip2.City) GeoInfoModel {
	gim := GeoInfoModel{
		IP:           IPString,
		CountryName:  unknown,
		Subdivisions: []string{unknown},
	}
	if record == nil {
		return gim
	}
	return gim.geoInfoRecord(IPString, record)
}

func (gi *GeoInfoModel) geoInfoRecord(IPString string, record *geoip2.City) GeoInfoModel {
	countryName := gi.countryName(record)
	subdivisions := gi.subdivisions(record)
	return GeoInfoModel{
		IP:           IPString,
		CountryName:  countryName,
		Subdivisions: subdivisions,
	}
}

func (gi *GeoInfoModel) countryName(record *geoip2.City) string {
	countryName := record.Country.Names["en"]
	if gi.emptyName(countryName) {
		countryName = unknown
	}
	return countryName
}

func (gi *GeoInfoModel) subdivisions(record *geoip2.City) []string {
	if gi.emptySubdivisions(record) {
		return []string{unknown}
	}

	return gi.extractSubdivisions(record)
}

func (gi *GeoInfoModel) emptyName(name string) bool {
	return len(name) == 0
}

func (gi *GeoInfoModel) emptySubdivisions(record *geoip2.City) bool {
	return len(record.Subdivisions) == 0
}

func (gi *GeoInfoModel) extractSubdivisions(record *geoip2.City) []string {
	subdivisions := []string{}

	for _, val := range record.Subdivisions {
		var subdivisionName = val.Names["en"]
		if gi.emptyName(subdivisionName) {
			subdivisionName = unknown
		}

		subdivisions = append(subdivisions, subdivisionName)
	}

	return subdivisions
}
