package geoinfo

import (
	"fmt"
	"testing"

	"github.com/oschwald/geoip2-golang"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	geoInfoRepoTest   *GeoInfoRepository
	geoIPDBReaderMock *mockGeoIPDBReader
)

type TSGeoInfoRepository struct{ suite.Suite }

func TestRunTSGeoInfo(t *testing.T) {
	suite.Run(t, new(TSGeoInfoRepository))
}

func (ts *TSGeoInfoRepository) BeforeTest(_, _ string) {
	geoIPDBReaderMock = newMockGeoIPDBReader(ts.T())
	geoInfoRepoTest = New("")
	geoInfoRepoTest.db = geoIPDBReaderMock
}

func (ts *TSGeoInfoRepository) TestSetDefaultDBPath() {
	ts.Equal(defaultDBPath, geoInfoRepoTest.dbPath)
}

func (ts *TSGeoInfoRepository) TestCloseCalled() {
	geoIPDBReaderMock.On("Close").Return(nil)
	err := geoInfoRepoTest.Close()

	ts.NoError(err)
	ts.True(geoIPDBReaderMock.AssertNumberOfCalls(ts.T(), "Close", 1))
}
func (ts *TSGeoInfoRepository) TestCloseCalledError() {
	geoIPDBReaderMock.On("Close").Return(fmt.Errorf("fakeError closing db"))
	err := geoInfoRepoTest.Close()

	ts.ErrorContains(err, "closing db")
}

func (ts *TSGeoInfoRepository) TestGetIPInfoError() {
	geoIPDBReaderMock.On("City", mock.Anything).Return(nil, fmt.Errorf("fake error on record"))

	record := geoInfoRepoTest.GetIPInfo("122.122.12.22")
	ts.Equal("122.122.12.22", record.IP())
	ts.Equal(unknown, record.CountryName())
	ts.Equal(unknown, record.Subdivisions()[0])
}
func (ts *TSGeoInfoRepository) TestGetIPInfo() {
	geoIPDBReaderMock.On("City", mock.Anything).Return(&geoip2.City{}, nil)

	record := geoInfoRepoTest.GetIPInfo("122.122.12.24")
	ts.Equal("122.122.12.24", record.IP())
}

func (ts *TSGeoInfoRepository) TestSetDBPath() {
	geoInfoRepoTest.setDBPath("otherFile.db")

	ts.Equal("otherFile.db", geoInfoRepoTest.dbPath)
}

func (ts *TSGeoInfoRepository) TestOpenDB() {
	err := geoInfoRepoTest.OpenDB()
	ts.ErrorContains(err, "error on opening the db")
}
