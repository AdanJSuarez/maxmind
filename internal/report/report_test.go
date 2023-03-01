package report

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	reportTest    *Report
	logParserMock logParser
	geoInfoMock   geoInfo
)

type TSReport struct{ suite.Suite }

func TestRunTSReport(t *testing.T) {
	suite.Run(t, new(TSReport))
}

func (ts *TSReport) BeforeTest(_, _ string) {
	linesCh := make(chan string, 10)
	logParserMock = newMockLogParser(ts.T())
	geoInfoMock = newMockGeoInfo(ts.T())
	reportTest = New(logParserMock, geoInfoMock, linesCh)
}

func (ts *TSReport) TestShouldExcludeTrue1() {
	logTrue1 := "/7b0744/css/vegguide-combined.css"
	actual := reportTest.shouldExclude(logTrue1)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue2() {
	logTrue2 := "/7b0744/css/vegguide-combined"
	actual := reportTest.shouldExclude(logTrue2)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse3() {
	logFalse3 := "/7b0744/csst/vegguide-combined"
	actual := reportTest.shouldExclude(logFalse3)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeFalse4() {
	logFalse4 := "/7b0744/cs/vegguide-combined"
	actual := reportTest.shouldExclude(logFalse4)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeTrue5() {
	logTrue5 := "/7b0744/cs/vegguide-combined.rss"
	actual := reportTest.shouldExclude(logTrue5)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue6() {
	logTrue6 := "/7b0744/cs/vegguide-combined.rss/"
	actual := reportTest.shouldExclude(logTrue6)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse7() {
	logFalse7 := "/7b0744/cs/vegguide-combined.rs"
	actual := reportTest.shouldExclude(logFalse7)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeFalse8() {
	logFalse8 := "/7b0744/cs/vegguide-combined.rsss"
	actual := reportTest.shouldExclude(logFalse8)
	ts.False(actual)
}
