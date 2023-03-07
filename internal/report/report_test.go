package report

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	reportTest *Report
)

type TSReport struct{ suite.Suite }

func TestRunTSReport(t *testing.T) {
	suite.Run(t, new(TSReport))
}

func (ts *TSReport) BeforeTest(_, _ string) {
	reportTest = New()
}

func (ts *TSReport) TestShouldExcludeTrue1() {
	logTrue := "/7b0744/css/vegguide-combined.css"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue2() {
	logTrue := "/7b0744/css/vegguide-combined"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse3() {
	logFalse := "/7b0744/csst/vegguide-combined"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeFalse4() {
	logFalse := "/7b0744/cs/vegguide-combined"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeTrue5() {
	logTrue := "/7b0744/cs/vegguide-combined.rss"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue6() {
	logTrue := "/7b0744/cs/vegguide-combined.rss/"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse7() {
	logFalse := "/7b0744/cs/vegguide-combined.rs"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeFalse8() {
	logFalse := "/7b0744/cs/vegguide-combined.rsss"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeTrue9() {
	logTrue := "/images/ratings/blue-3-00.png"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue10() {
	logTrue := "/7b0744/images/ratings/green-0-00.png"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse11() {
	logFalse := "/image/ratings/blue-3-00.png"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeFalse12() {
	logFalse := "images/ratings/blue-3-00.png"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeFalse13() {
	logFalse := "/imagess/ratings/blue-3-00.png"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}
