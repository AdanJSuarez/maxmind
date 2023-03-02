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
	countriesTest = NewCountries()
}

func (ts *TSCountries) TestCountriesInitialization() {
	ts.NotNil(countriesTest.countries)
}

func (ts *TSCountries) TestAddToCountries1() {
	countriesTest.AddToCountries("United States", "Montana", "/")

	country := countriesTest.countries["United States"]
	ts.Equal(int64(1), country.counter)

	subdivision := country.subdivisions["Montana"]
	ts.Equal(int64(1), subdivision.counter)

	webpage := subdivision.webpages["/"]
	ts.Equal(int64(1), webpage)
}

func (ts *TSCountries) TestAddToCountries2() {
	countriesTest.AddToCountries("United States", "Montana", "/")
	countriesTest.AddToCountries("United States", "Montana", "/")

	country := countriesTest.countries["United States"]
	ts.Equal(int64(2), country.counter)

	subdivision := country.subdivisions["Montana"]
	ts.Equal(int64(2), subdivision.counter)

	webpage := subdivision.webpages["/"]
	ts.Equal(int64(2), webpage)
}

func (ts *TSCountries) TestAddToCountries3() {
	countriesTest.AddToCountries("United States", "Montana", "/")
	countriesTest.AddToCountries("United States", "Montana", "/turbo")

	country := countriesTest.countries["United States"]
	ts.Equal(int64(2), country.counter)

	subdivision := country.subdivisions["Montana"]
	ts.Equal(int64(2), subdivision.counter)

	webpage1 := subdivision.webpages["/"]
	ts.Equal(int64(1), webpage1)

	webpage2 := subdivision.webpages["/turbo"]
	ts.Equal(int64(1), webpage2)
}

func (ts *TSCountries) TestAddToCountries4() {
	countriesTest.AddToCountries("United States", "Alabama", "/")
	countriesTest.AddToCountries("United States", "Montana", "/turbo")

	country := countriesTest.countries["United States"]
	ts.Equal(int64(2), country.counter)

	subdivision1 := country.subdivisions["Montana"]
	ts.Equal(int64(1), subdivision1.counter)

	webpage1 := subdivision1.webpages["/"]
	ts.Equal(int64(0), webpage1)

	webpage2 := subdivision1.webpages["/turbo"]
	ts.Equal(int64(1), webpage2)

	subdivision2 := country.subdivisions["Alabama"]
	ts.Equal(int64(1), subdivision1.counter)

	webpage3 := subdivision2.webpages["/"]
	ts.Equal(int64(1), webpage3)

	webpage4 := subdivision2.webpages["/turbo"]
	ts.Equal(int64(0), webpage4)
}
