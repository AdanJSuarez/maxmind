package countries

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var infoTest Info

type TSInfo struct{ suite.Suite }

func TestRunTSInfo(t *testing.T) {
	suite.Run(t, new(TSInfo))
}

func (ts *TSInfo) BeforeTest(_, _ string) {
	infoTest = NewInfo("Canada", "/", 1)
}

func (ts *TSInfo) TestGetters() {
	ts.Equal("Canada", infoTest.Name())
	ts.Equal(int64(1), infoTest.Visit())
	ts.Equal("/", infoTest.TopPage())
}
