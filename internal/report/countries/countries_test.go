package countries

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	countriesTest *Countries
)

type TSCountries struct{ suite.Suite }

func TestRunTSCountries(t *testing.T) {
	suite.Run(t, new(TSCountries))
}

func (ts *TSCountries) BeforeTest(_, _ string) {
	countriesTest = New()
}

func (ts *TSCountries) TestCountriesInitialization() {
	ts.NotNil(countriesTest.countries)
}

func (ts *TSCountries) TestAddToCountries1() {
	countriesTest.AddToCountries("United States", "Montana", "/")

	country := countriesTest.countries.Children()["United States"]
	ts.Equal(int64(1), country.Counter())

	subdivision := country.Children()["Montana"]
	ts.Equal(int64(1), subdivision.Counter())

	webpage := subdivision.Data()["/"]
	ts.Equal(int64(1), webpage)
}

func (ts *TSCountries) TestAddToCountries2() {
	countriesTest.AddToCountries("United States", "Montana", "/")
	countriesTest.AddToCountries("United States", "Montana", "/")

	country := countriesTest.countries.Children()["United States"]
	ts.Equal(int64(2), country.Counter())

	subdivision := country.Children()["Montana"]
	ts.Equal(int64(2), subdivision.Counter())

	webpage := subdivision.Data()["/"]
	ts.Equal(int64(2), webpage)
}

func (ts *TSCountries) TestAddToCountries3() {
	countriesTest.AddToCountries("United States", "Montana", "/")
	countriesTest.AddToCountries("United States", "Montana", "/turbo")

	country := countriesTest.countries.Children()["United States"]
	ts.Equal(int64(2), country.Counter())

	subdivision := country.Children()["Montana"]
	ts.Equal(int64(2), subdivision.Counter())

	webpage1 := subdivision.Data()["/"]
	ts.Equal(int64(1), webpage1)

	webpage2 := subdivision.Data()["/turbo"]
	ts.Equal(int64(1), webpage2)
}

func (ts *TSCountries) TestAddToCountries4() {
	countriesTest.AddToCountries("United States", "Alabama", "/")
	countriesTest.AddToCountries("United States", "Montana", "/turbo")

	country := countriesTest.countries.Children()["United States"]
	ts.Equal(int64(2), country.Counter())

	subdivision1 := country.Children()["Montana"]
	ts.Equal(int64(1), subdivision1.Counter())

	webpage1 := subdivision1.Data()["/"]
	ts.Equal(int64(0), webpage1)

	webpage2 := subdivision1.Data()["/turbo"]
	ts.Equal(int64(1), webpage2)

	subdivision2 := country.Children()["Alabama"]
	ts.Equal(int64(1), subdivision1.Counter())

	webpage3 := subdivision2.Data()["/"]
	ts.Equal(int64(1), webpage3)

	webpage4 := subdivision2.Data()["/turbo"]
	ts.Equal(int64(0), webpage4)
}
