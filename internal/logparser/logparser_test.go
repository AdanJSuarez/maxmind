package logparser

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	logTest1           = `183.60.212.148 - - [26/Aug/2014:06:26:39 -0600] "GET /entry/15205 HTTP/1.1" 200 4865 "-" "Mozilla/5.0 (compatible; EasouSpider; +http://www.easou.com/search/spider.html)"`
	logTest2           = `68.180.225.35 - - [26/Aug/2014:06:44:30 -0600] "GET /entry/20153 HTTP/1.1" 200 4539 "-" "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)"`
	logIncompleteTest3 = `23.239.8.38 - - [26/Aug/2014:06:59:31 -0600] "GET / HTTP/1.1" 301 178`
	logCorruptTest4    = `😀😁`
)

var (
	logParserTest *LogParser
)

type TSLogParser struct{ suite.Suite }

func TestRunTSLogParser(t *testing.T) {
	suite.Run(t, new(TSLogParser))
}

func (ts *TSLogParser) BeforeTest(_, _ string) {
	var err error
	logParserTest, err = New()
	ts.NoError(err)
}

func (ts *TSLogParser) TestLogLine1() {
	actual, err := logParserTest.Parse(logTest1)
	ts.NoError(err)
	ts.Equal("183.60.212.148", actual.IP)
	ts.Equal("26/Aug/2014:06:26:39 -0600", actual.TS)
	ts.Equal("GET", actual.RequestMethod)
	ts.Equal("/entry/15205", actual.RequestPath)
	ts.Equal(int64(200), actual.StatusCode)
	ts.Equal(int64(4865), actual.Size)
}

func (ts *TSLogParser) TestLogLine2() {
	actual, err := logParserTest.Parse(logTest2)
	ts.NoError(err)
	ts.Equal("68.180.225.35", actual.IP)
	ts.Equal("26/Aug/2014:06:44:30 -0600", actual.TS)
	ts.Equal("GET", actual.RequestMethod)
	ts.Equal("/entry/20153", actual.RequestPath)
	ts.Equal(int64(200), actual.StatusCode)
	ts.Equal(int64(4539), actual.Size)
}

func (ts *TSLogParser) TestLogLine3() {
	actual, err := logParserTest.Parse(logIncompleteTest3)
	ts.ErrorContains(err, "does not have all the matches")
	ts.Empty(actual)
}
func (ts *TSLogParser) TestLogLine4() {
	actual, err := logParserTest.Parse(logCorruptTest4)
	ts.ErrorContains(err, "does not have all the matches")
	ts.Empty(actual)
}

func (ts *TSLogParser) TestParseValidInteger() {
	actual := logParserTest.parseStringToInt64("333")
	ts.Equal(int64(333), actual)
}

func (ts *TSLogParser) TestParseInvalidInteger() {
	actual := logParserTest.parseStringToInt64("turbo")
	ts.Equal(int64(0), actual)
}
