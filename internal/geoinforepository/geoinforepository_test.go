package geoinfo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	geoInfoRepoTest   *GeoInfoRepository
	geoIPDBReaderMock *mockGeoIPReader
)

type TSGeoInfoRepository struct{ suite.Suite }

func TestRunTSGeoInfo(t *testing.T) {
	suite.Run(t, new(TSGeoInfoRepository))
}

func (ts *TSGeoInfoRepository) BeforeTest(_, _ string) {
	geoIPDBReaderMock = newMockGeoIPReader(ts.T())
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
	ts.Empty(record)
}

func (ts *TSGeoInfoRepository) TestSetDBPath() {
	geoInfoRepoTest.setDBPath("otherFile.db")

	ts.Equal("otherFile.db", geoInfoRepoTest.dbPath)
}
