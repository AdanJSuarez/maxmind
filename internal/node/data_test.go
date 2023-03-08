package node

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var dataTest Data

type TSData struct{ suite.Suite }

func TestRunTSData(t *testing.T) {
	suite.Run(t, new(TSData))
}

func (ts *TSData) BeforeTest(_, _ string) {
	dataTest = NewData("Spain", 3)
}

func (ts *TSData) TestData() {
	ts.Equal("Spain", dataTest.Name())
	ts.Equal(int64(3), dataTest.Counter())
}
