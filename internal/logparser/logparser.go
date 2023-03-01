package logparser

import (
	"log"
	"regexp"
	"strconv"
)

// LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-agent}i\"" combined
// LogExample: 183.60.212.148 - - [26/Aug/2014:06:26:39 -0600] "GET /entry/15205 HTTP/1.1" 200 4865 "-" "Mozilla/5.0 (compatible; EasouSpider; +http://www.easou.com/search/spider.html)"
// logsFormat = `\[$time_stamp\] \"$http_method $request_path $_\" $response_code - $_ $_ $_ - \"$ips\" \"$_\" \"$_\" \"$_\" \"$_\"`
const (
	logsFormat = `$ip $_ $_ \[$time_stamp\] \"$request_method $request_path $protocol\" $status_code $size \"$_\" \"$_ http$web_page\.html"`
)

type Log struct {
	IP            string
	TS            string
	RequestMethod string
	RequestPath   string
	StatusCode    int64
	Size          int64
	WebPage       string
}

type LogParser struct {
	regex *regexp.Regexp
}

func New() *LogParser {
	regexFormat := regexp.MustCompile(`\$([\w_]*)`).ReplaceAllString(logsFormat, `(?P<$1>.*)`)
	regex, err := regexp.Compile(regexFormat)
	if err != nil {
		log.Println("error: ", err)
	}
	return &LogParser{
		regex: regex,
	}
}

func (lp *LogParser) Parse(line string) Log {
	matches := lp.regex.FindStringSubmatch(line)
	return Log{
		IP:            matches[1],
		TS:            matches[4],
		RequestMethod: matches[5],
		RequestPath:   matches[6],
		StatusCode:    lp.parseStringToInt64(matches[8]),
		Size:          lp.parseStringToInt64(matches[9]),
		WebPage:       matches[14],
	}
}

func (lp *LogParser) parseStringToInt64(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		log.Printf("error parsing string: %s. An empty value is assigned: %v\n", s, err)
	}
	return result
}
