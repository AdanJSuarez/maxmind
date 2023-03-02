package countries

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	subdivisionTest *Subdivision
)

type TSSubdivision struct{ suite.Suite }

func TestRunTSSubdivision(t *testing.T) {
	suite.Run(t, new(TSSubdivision))
}

func (ts *TSSubdivision) BeforeTest(_, _ string) {
	subdivisionTest = newSubdivision("Ontario", "/")
}

func (ts *TSSubdivision) TestSubdivisionInitialized() {
	ts.NotNil(subdivisionTest.webpages)
	ts.Equal(subdivisionTest.counter, int64(1))
}

func (ts *TSSubdivision) TestAddToSubdivisionExistingPage() {
	subdivisionTest.addToSubdivision("/")
	ts.Equal(int64(2), subdivisionTest.counter)
	ts.Equal(int64(2), subdivisionTest.webpages["/"])
}

func (ts *TSSubdivision) TestAddToSubdivisionNotExistingPage() {
	subdivisionTest.addToSubdivision("/turbo")
	ts.Equal(int64(2), subdivisionTest.counter)
	ts.Equal(int64(1), subdivisionTest.webpages["/"])
	ts.Equal(int64(1), subdivisionTest.webpages["/turbo"])
}
