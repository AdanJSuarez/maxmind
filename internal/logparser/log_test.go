package logparser

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var logTest Log

type TSLog struct{ suite.Suite }

func TestRunTSLog(t *testing.T) {
	suite.Run(t, new(TSLog))
}

func (ts *TSLog) BeforeTest(_, _ string) {
	logTest = NewLog("123.123.55.55", "2020/03/11", "GET", "/turbo", int64(200), 1)
}

func (ts *TSLog) TestInitializedLog() {
	ts.Equal("123.123.55.55", logTest.IP())
	ts.Equal("2020/03/11", logTest.TS())
	ts.Equal("GET", logTest.RequestMethod())
	ts.Equal("/turbo", logTest.RequestPath())
	ts.Equal(int64(200), logTest.StatusCode())
	ts.Equal(int64(1), logTest.Size())
}
