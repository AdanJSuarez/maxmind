package logparser

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	logTest = `183.60.212.148 - - [26/Aug/2014:06:26:39 -0600] "GET /entry/15205 HTTP/1.1" 200 4865 "-" "Mozilla/5.0 (compatible; EasouSpider; +http://www.easou.com/search/spider.html)"`
)

var (
	parserTest *LogParser
)

type TSLogParser struct{ suite.Suite }

func TestRunTSLogParser(t *testing.T) {
	suite.Run(t, new(TSLogParser))
}

func (ts *TSLogParser) BeforeTest(_, _ string) {
	parserTest = New()
}

func (ts *TSLogParser) TestLogLine() {
	actual := parserTest.Parse(logTest)
	ts.Equal("183.60.212.148", actual.IP)
	ts.Equal("26/Aug/2014:06:26:39 -0600", actual.TS)
	ts.Equal("GET", actual.RequestMethod)
	ts.Equal("/entry/15205", actual.RequestPath)
	ts.Equal(int64(200), actual.StatusCode)
	ts.Equal(int64(4865), actual.Size)
}

func (ts *TSLogParser) TestParseValidInteger() {
	actual := parserTest.parseStringToInt64("333")
	ts.Equal(int64(333), actual)
}

func (ts *TSLogParser) TestParseInvalidInteger() {
	actual := parserTest.parseStringToInt64("turbo")
	ts.Equal(int64(0), actual)
}
