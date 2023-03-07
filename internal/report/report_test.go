package report

import (
	"testing"

	"github.com/AdanJSuarez/maxmind/internal/report/countries"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	reportTest *Report
	dataMock   *mockCountriesData
)

type TSReport struct{ suite.Suite }

func TestRunTSReport(t *testing.T) {
	suite.Run(t, new(TSReport))
}

func (ts *TSReport) BeforeTest(_, _ string) {
	dataMock = newMockCountriesData(ts.T())
	reportTest = New()
	reportTest.data = dataMock
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
func (ts *TSReport) TestShouldExcludeFalse14() {
	logFalse := "/js/ratings"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}
func (ts *TSReport) TestShouldExcludeTrue15() {
	logTrue := "aff09/js/ratings"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue16() {
	logTrue := "/js/ratings/value.js"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue17() {
	logTrue := "/entry-images/ratings/value"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue18() {
	logTrue := "/lsfj/entry-images/ratings/value"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse19() {
	logFalse := "/entry-image/ratings"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeTrue20() {
	logTrue := "/lsfj/static/ratings/value"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}
func (ts *TSReport) TestShouldExcludeTrue21() {
	logTrue := "/static/ratings/value"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse22() {
	logFalse := "/statics/ratings"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeTrue22() {
	logTrue := "/turbo/robots.txt"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue23() {
	logTrue := "/turbo/robots.txt/"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue24() {
	logTrue := "/robots.txt"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse25() {
	logFalse := "/robots.tx"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeTrue26() {
	logTrue := "/favicon.ico"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue27() {
	logTrue := "lasdjfls/lsdkjf/favicon.ico"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue28() {
	logTrue := "/favicon.ico/"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse29() {
	logFalse := "/favicon.ic"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeTrue30() {
	logTrue := "/turbo.rss"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue31() {
	logTrue := "/turbo.rss/"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue32() {
	logTrue := "/ldjfs/turbo.rss"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeFalse33() {
	logFalse := "/turbo.rs"
	actual := reportTest.ShouldExclude(logFalse)
	ts.False(actual)
}

func (ts *TSReport) TestShouldExcludeTrue34() {
	logTrue := "/turbo.atom"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue35() {
	logTrue := "/turbo.atom/"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestShouldExcludeTrue36() {
	logTrue := "ldjf/ldjfii/turbo.atom"
	actual := reportTest.ShouldExclude(logTrue)
	ts.True(actual)
}

func (ts *TSReport) TestSubdivision() {
	subdivisions := []string{"Madrid", "Alcorcon"}
	actual := reportTest.Subdivision(subdivisions)
	ts.Equal("Madrid", actual)
}

func (ts *TSReport) TestAddData() {
	dataMock.On("AddToCountries", mock.Anything, mock.Anything, mock.Anything).Return()

	reportTest.AddData("Spain", "Tenerife", "/turbo")
	ts.True(dataMock.AssertNumberOfCalls(ts.T(), "AddToCountries", 1))
}

func (ts *TSReport) TestPrintReport() {
	infos := []countries.Info{
		{Name: "Spain", Visit: 3, TopPage: "/turbo"},
		{Name: "Canada", Visit: 2, TopPage: "/rambo"},
	}
	dataMock.On("TopAreas", mock.Anything, mock.Anything, mock.Anything).Return(infos)
	dataMock.On("Name").Return("Countries")

	reportTest.Generate()
	ts.True(dataMock.AssertNumberOfCalls(ts.T(), "TopAreas", 2))
	ts.True(dataMock.AssertNumberOfCalls(ts.T(), "Name", 1))
}
