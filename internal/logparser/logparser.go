package logparser

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	minimumMatchesLength = 10
	integerBase          = 10
	logsFormat           = `$ip $_ $_ \[$time_stamp\] \"$request_method $request_path $protocol\" $status_code $size \"$_\" \"$_\"`
	ipIdx                = 1
	timeStampIdx         = 4
	requestMethodIdx     = 5
	requestPathIdx       = 6
	statusCodeIdx        = 8
	sizeIdx              = 9
)

type LogParser struct {
	regex *regexp.Regexp
}

// New returns a initialized instance of LogParser. An error otherwise.
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

// Parse reads the line and extract the Log from it. It returns
// an error otherwise.
func (lp *LogParser) Parse(line string) (Log, error) {
	matches := lp.regex.FindStringSubmatch(line)
	if !lp.hasAllNeededMatches(matches) {
		return Log{}, fmt.Errorf("does not have the matches")
	}

	log := NewLog(
		matches[ipIdx],
		matches[timeStampIdx],
		matches[requestMethodIdx],
		matches[requestPathIdx],
		lp.parseStringToInt64(matches[statusCodeIdx]),
		lp.parseStringToInt64(matches[sizeIdx]),
	)

	return log, nil
}

func (lp *LogParser) parseStringToInt64(s string) int64 {
	result, err := strconv.ParseInt(s, integerBase, 0)
	if err != nil {
		fmt.Printf("error parsing string: %s. An empty value is assigned: %v\n", s, err)
	}

	return result
}

func (lp *LogParser) hasAllNeededMatches(matches []string) bool {
	return len(matches) >= minimumMatchesLength
}
