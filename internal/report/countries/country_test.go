package countries

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	countryTest *Country
)

type TSCountry struct{ suite.Suite }

func TestRunTSCountry(t *testing.T) {
	suite.Run(t, new(TSCountry))
}

func (ts *TSCountry) BeforeTest(_, _ string) {
	countryTest = newCountry("Ontario", "/")
}

func (ts *TSCountry) TestCountryInitialized() {
	ts.NotNil(countryTest.subdivisions)
	ts.Equal(int64(1), countryTest.counter)

	subdivision := countryTest.subdivisions["Ontario"]
	ts.Equal(int64(1), subdivision.counter)
}
func (ts *TSCountry) TestAddToCountry1() {
	countryTest.addToCountry("Ontario", "/")
	ts.Equal(int64(2), countryTest.counter)

	subdivision := countryTest.subdivisions["Ontario"]
	ts.Equal(int64(2), subdivision.counter)
}

func (ts *TSCountry) TestAddToCountry2() {
	countryTest.addToCountry("Ontario", "/turbo")
	ts.Equal(int64(2), countryTest.counter)

	subdivision := countryTest.subdivisions["Ontario"]
	ts.Equal(int64(2), subdivision.counter)
}

func (ts *TSCountry) TestAddToCountry3() {
	countryTest.addToCountry("Montana", "/")
	ts.Equal(int64(2), countryTest.counter)

	subdivisionOntario := countryTest.subdivisions["Ontario"]
	ts.Equal(int64(1), subdivisionOntario.counter)

	subdivisionMontana := countryTest.subdivisions["Montana"]
	ts.Equal(int64(1), subdivisionMontana.counter)
}

func (ts *TSCountry) TestAddToCountry4() {
	countryTest.addToCountry("Montana", "/")
	ts.Equal(int64(2), countryTest.counter)

	countryTest.addToCountry("Montana", "/turbo")
	ts.Equal(int64(3), countryTest.counter)

	subdivisionOntario := countryTest.subdivisions["Ontario"]
	ts.Equal(int64(1), subdivisionOntario.counter)

	subdivisionMontana := countryTest.subdivisions["Montana"]
	ts.Equal(int64(2), subdivisionMontana.counter)
}
