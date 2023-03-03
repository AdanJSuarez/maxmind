package geoinfo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var geoInfoModelTest GeoInfoModel

type TSGeoInfoModel struct{ suite.Suite }

func TestRunTSGeoInfoModel(t *testing.T) {
	suite.Run(t, new(TSGeoInfoModel))
}

func (ts *TSGeoInfoModel) BeforeTest(_, _ string) {
	geoInfoModelTest = newGeoInfoModel("", nil)
}

func (ts *TSGeoInfoModel) Test() {}
