package geoinfo

import (
	"testing"

	"github.com/oschwald/geoip2-golang"
	"github.com/stretchr/testify/suite"
)

var (
	geoInfoModelTest GeoInfoData
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
	recordTest2 = geoip2.City{
		Country: struct {
			GeoNameID         uint              `maxminddb:"geoname_id"`
			IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
			IsoCode           string            `maxminddb:"iso_code"`
			Names             map[string]string `maxminddb:"names"`
		}{
			Names: map[string]string{},
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
	recordTest3 = geoip2.City{
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
				Names: map[string]string{},
			},
		},
	}
	recordTest4 = geoip2.City{
		Country: struct {
			GeoNameID         uint              `maxminddb:"geoname_id"`
			IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
			IsoCode           string            `maxminddb:"iso_code"`
			Names             map[string]string `maxminddb:"names"`
		}{
			Names: nil,
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
				Names: nil,
			},
		},
	}
)

type TSGeoInfoModel struct{ suite.Suite }

func TestRunTSGeoInfoModel(t *testing.T) {
	suite.Run(t, new(TSGeoInfoModel))
}

func (ts *TSGeoInfoModel) BeforeTest(_, _ string) {
	geoInfoModelTest = newGeoInfoData("122.122.12.22", nil)
}

func (ts *TSGeoInfoModel) TestGeoInfoModelInitializedNil() {
	ts.Equal("122.122.12.22", geoInfoModelTest.IP())
	ts.Equal(unknown, geoInfoModelTest.CountryName())
	ts.Equal(unknown, geoInfoModelTest.Subdivisions()[0])
}

func (ts *TSGeoInfoModel) TestGeoInfoModelCompleteRecord() {
	geoInfoModelTest = newGeoInfoData("122.122.12.23", &recordTest1)
	ts.Equal("122.122.12.23", geoInfoModelTest.IP())
	ts.Equal("Spain", geoInfoModelTest.CountryName())
	ts.Equal("Tenerife", geoInfoModelTest.Subdivisions()[0])
}

func (ts *TSGeoInfoModel) TestGeoInfoModelNoCountryName() {
	geoInfoModelTest = newGeoInfoData("122.122.12.24", &recordTest2)
	ts.Equal("122.122.12.24", geoInfoModelTest.IP())
	ts.Equal(unknown, geoInfoModelTest.CountryName())
	ts.Equal("Tenerife", geoInfoModelTest.Subdivisions()[0])
}

func (ts *TSGeoInfoModel) TestGeoInfoModelNoSubdivisions() {
	geoInfoModelTest = newGeoInfoData("122.122.12.24", &recordTest3)
	ts.Equal("122.122.12.24", geoInfoModelTest.IP())
	ts.Equal("Spain", geoInfoModelTest.CountryName())
	ts.Equal(unknown, geoInfoModelTest.Subdivisions()[0])
}

func (ts *TSGeoInfoModel) TestGeoInfoModelNilMaps() {
	geoInfoModelTest = newGeoInfoData("122.122.12.25", &recordTest4)
	ts.Equal("122.122.12.25", geoInfoModelTest.IP())
	ts.Equal(unknown, geoInfoModelTest.CountryName())
	ts.Equal(unknown, geoInfoModelTest.Subdivisions()[0])
}
