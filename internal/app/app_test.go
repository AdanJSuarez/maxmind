package app

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/AdanJSuarez/maxmind/internal/configuration"
	"github.com/AdanJSuarez/maxmind/internal/geoinfo"
	"github.com/AdanJSuarez/maxmind/internal/logparser"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	appTest           *App
	logReaderMock     *mockLogReader
	geoInfoMock       *mockGeoInfo
	logParserMock     *mockLogParser
	countryReportMock *mockCountryReport
	wg                sync.WaitGroup
)

type TSApp struct{ suite.Suite }

func TestRunTSApp(t *testing.T) {
	suite.Run(t, new(TSApp))
}

func (ts *TSApp) BeforeTest(_, _ string) {
	logReaderMock = newMockLogReader(ts.T())
	geoInfoMock = newMockGeoInfo(ts.T())
	logParserMock = newMockLogParser(ts.T())
	countryReportMock = newMockCountryReport(ts.T())

	appTest = &App{
		wg:      &wg,
		config:  configuration.Configuration{},
		linesCh: make(chan string, 10),
	}
	appTest.logReader = logReaderMock
	appTest.geoInfo = geoInfoMock
	appTest.logParser = logParserMock
	appTest.report = countryReportMock
}

func (ts *TSApp) TestInitializeApp() {
	geoInfoMock.On("OpenDB").Return(nil)
	logReaderMock.On("Open").Return(nil)
	err := appTest.initializeApp(appTest.wg)
	ts.NoError(err)
}
func (ts *TSApp) TestInitializeAppOnError() {
	logReaderMock.On("Open").Return(fmt.Errorf("fakeError2"))
	err := appTest.initializeApp(appTest.wg)
	ts.Error(err)
}

func (ts *TSApp) TestInitializeAppOnError2() {
	logReaderMock.On("Open").Return(nil)
	geoInfoMock.On("OpenDB").Return(fmt.Errorf("fakeError3"))
	err := appTest.initializeApp(appTest.wg)
	ts.Error(err)
}
func (ts *TSApp) TestInitializationOnError() {
	appTest, err := New(&wg, configuration.Configuration{})
	ts.Nil(appTest)
	ts.Error(err)
}

func (ts *TSApp) TestStart() {
	logReaderMock.On("ReadLinesFromFile").Return()
	logParserMock.On("Parse", mock.Anything).Return(logparser.Log{RequestPath: "/turbo"}, nil)
	geoInfoMock.On("GetIPInfo", mock.Anything).Return(geoinfo.GeoInfoData{})
	countryReportMock.On("ShouldExclude", mock.Anything).Return(false)
	countryReportMock.On("Subdivision", mock.Anything).Return("fake")
	countryReportMock.On("AddData", mock.Anything, mock.Anything, mock.Anything).Return()

	appTest.linesCh <- "fakelog"
	go appTest.Start()
	time.Sleep(2 * time.Second)

	ts.True(logReaderMock.AssertNumberOfCalls(ts.T(), "ReadLinesFromFile", 1))
	ts.True(logParserMock.AssertNumberOfCalls(ts.T(), "Parse", 1))
	ts.True(countryReportMock.AssertNumberOfCalls(ts.T(), "ShouldExclude", 1))
	ts.True(geoInfoMock.AssertNumberOfCalls(ts.T(), "GetIPInfo", 1))
	ts.True(countryReportMock.AssertNumberOfCalls(ts.T(), "AddData", 1))
	ts.True(countryReportMock.AssertNumberOfCalls(ts.T(), "Subdivision", 1))
}

func (ts *TSApp) TestCloseOnError() {
	geoInfoMock.On("Close", mock.Anything).Return(fmt.Errorf("fakeErrorGeo"))
	logReaderMock.On("Close").Return(fmt.Errorf("fakeErrorReader"))
	ts.NotPanics(func() { appTest.Close() })
}

func (ts *TSApp) TestCloseNoError() {
	geoInfoMock.On("Close", mock.Anything).Return(nil)
	logReaderMock.On("Close").Return(nil)
	ts.NotPanics(func() { appTest.Close() })
}

func (ts *TSApp) TestPopulateData() {
	logParserMock.On("Parse", mock.Anything).Return(logparser.Log{RequestPath: "/turbo"}, nil)
	countryReportMock.On("ShouldExclude", mock.Anything).Return(true)

	appTest.linesCh <- "fakelog"
	appTest.wg.Add(1)
	go appTest.populateData()
	time.Sleep(time.Second)
	close(appTest.linesCh)

	ts.True(logParserMock.AssertNumberOfCalls(ts.T(), "Parse", 1))
	ts.True(countryReportMock.AssertNumberOfCalls(ts.T(), "ShouldExclude", 1))
}
