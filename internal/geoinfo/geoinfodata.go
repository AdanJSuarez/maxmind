package geoinfo

import "github.com/oschwald/geoip2-golang"

const (
	english = "en"
	unknown = "unknown"
)

// GeoInfoData represents the information about a specific IP.
// This abstract the the API client we use.
type GeoInfoData struct {
	ip           string
	countryName  string
	subdivisions []string
}

// newGeoInfoData returns a initialized instance of GeoInfoData.
func newGeoInfoData(IPString string, record *geoip2.City) GeoInfoData {
	gim := GeoInfoData{
		ip:           IPString,
		countryName:  unknown,
		subdivisions: []string{unknown},
	}
	if record == nil {
		return gim
	}
	return gim.geoInfoData(IPString, record)
}

// Getters

func (gi *GeoInfoData) IP() string {
	return gi.ip
}

func (gi *GeoInfoData) CountryName() string {
	return gi.countryName
}

func (gi *GeoInfoData) Subdivisions() []string {
	return gi.subdivisions
}

func (gi *GeoInfoData) geoInfoData(IPString string, record *geoip2.City) GeoInfoData {
	countryName := gi.countryNameEnglish(record)
	subdivisions := gi.subdivisionsEnglish(record)
	return GeoInfoData{
		ip:           IPString,
		countryName:  countryName,
		subdivisions: subdivisions,
	}
}

func (gi *GeoInfoData) countryNameEnglish(record *geoip2.City) string {
	countryName := record.Country.Names[english]
	if gi.emptyName(countryName) {
		countryName = unknown
	}
	return countryName
}

func (gi *GeoInfoData) subdivisionsEnglish(record *geoip2.City) []string {
	if gi.emptySubdivisions(record) {
		return []string{unknown}
	}

	return gi.extractSubdivisions(record)
}

func (gi *GeoInfoData) emptyName(name string) bool {
	return len(name) == 0
}

func (gi *GeoInfoData) emptySubdivisions(record *geoip2.City) bool {
	return len(record.Subdivisions) == 0
}

func (gi *GeoInfoData) extractSubdivisions(record *geoip2.City) []string {
	subdivisions := []string{}

	for _, val := range record.Subdivisions {
		var subdivisionName = val.Names[english]
		if gi.emptyName(subdivisionName) {
			subdivisionName = unknown
		}

		subdivisions = append(subdivisions, subdivisionName)
	}

	return subdivisions
}
