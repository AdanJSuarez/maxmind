package geoinfo

import (
	"testing"

	"github.com/oschwald/geoip2-golang"
	"github.com/stretchr/testify/suite"
)

var (
	geoInfoModelTest GeoInfoModel
	recordTest1      = geoip2.City{
		Country: struct {
			GeoNameID         uint              `maxminddb:"geoname_id"`
			IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
			IsoCode           string            `maxminddb:"iso_code"`
			Names             map[string]string `maxminddb:"names"`
		}{
			Names: map[string]string{"en": "Spain"},
		},
		Subdivisions: []struct {
			GeoNameID uint              `maxminddb:"geoname_id"`
			IsoCode   string            `maxminddb:"iso_code"`
			Names     map[string]string `maxminddb:"names"`
		}{
			struct {
				GeoNameID uint              "maxminddb:\"geoname_id\""
				IsoCode   string            "maxminddb:\"iso_code\""
				Names     map[string]string "maxminddb:\"names\""
			}{
				Names: map[string]string{"en": "Tenerife"},
			},
		},
	}
)

type TSGeoInfoModel struct{ suite.Suite }

func TestRunTSGeoInfoModel(t *testing.T) {
	suite.Run(t, new(TSGeoInfoModel))
}

func (ts *TSGeoInfoModel) BeforeTest(_, _ string) {
	geoInfoModelTest = newGeoInfoModel("122.122.12.22", nil)
}

func (ts *TSGeoInfoModel) TestGeoInfoModelInitialized() {
	ts.Equal("122.122.12.22", geoInfoModelTest.IP)
	ts.Equal(unknown, geoInfoModelTest.CountryName)
	ts.Equal(unknown, geoInfoModelTest.Subdivisions[0])
}

func (ts *TSGeoInfoModel) TestGeoInfoModelInitialized2() {
	geoInfoModelTest = newGeoInfoModel("122.122.12.23", &recordTest1)
	ts.Equal("122.122.12.23", geoInfoModelTest.IP)
	ts.Equal("Spain", geoInfoModelTest.CountryName)
	ts.Equal("Tenerife", geoInfoModelTest.Subdivisions[0])
}
