package logparser

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// LogExample: 183.60.212.148 - - [26/Aug/2014:06:26:39 -0600] "GET /entry/15205 HTTP/1.1" 200 4865 "-" "Mozilla/5.0 (compatible; EasouSpider; +http://www.easou.com/search/spider.html)"

const (
	minimumMatchesLength = 10
	logsFormat           = `$ip $_ $_ \[$time_stamp\] \"$request_method $request_path $protocol\" $status_code $size \"$_\" \"$_\"`
)

type Log struct {
	IP            string
	TS            string
	RequestMethod string
	RequestPath   string
	StatusCode    int64
	Size          int64
}

type LogParser struct {
	regex *regexp.Regexp
}

func New() (*LogParser, error) {
	regexFormat := regexp.MustCompile(`\$([\w_]*)`).ReplaceAllString(logsFormat, `(?P<$1>.*)`)
	regex, err := regexp.Compile(regexFormat)
	if err != nil {
		return nil, err
	}

	logParser := &LogParser{
		regex: regex,
	}

	return logParser, nil
}

func (lp *LogParser) Parse(line string) (Log, error) {
	matches := lp.regex.FindStringSubmatch(line)
	if !lp.hasAllNeededMatches(matches) {
		return Log{}, fmt.Errorf("does not have all the matches: %s", line)
	}
	log := Log{
		IP:            matches[1],
		TS:            matches[4],
		RequestMethod: matches[5],
		RequestPath:   matches[6],
		StatusCode:    lp.parseStringToInt64(matches[8]),
		Size:          lp.parseStringToInt64(matches[9]),
	}
	return log, nil
}

func (lp *LogParser) parseStringToInt64(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		log.Printf("error parsing string: %s. An empty value is assigned: %v\n", s, err)
	}
	return result
}

func (lp *LogParser) hasAllNeededMatches(matches []string) bool {
	return len(matches) >= minimumMatchesLength
}
