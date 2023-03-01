package report

import (
	"log"
	"regexp"
)

const (
	cssOr         = `[a-f0-9]+/css/|`
	imagesOr      = `[a-f0-9]+/images/|/images/|`
	entryImagesOr = `/entry-images/|`
	staticOr      = `/static/|`
	robotstxtOr   = `/robots.txt/?$|`
	faviconicoOr  = `/favicon.ico/?$|`
	rssOr         = `\w\.rss/?$|`
	atomOr        = `\w\.atom/?$`
	logPatter     = cssOr + imagesOr + entryImagesOr + staticOr + robotstxtOr + faviconicoOr + rssOr + atomOr
)

type Report struct {
	logParser logParser
	geoInfo   geoInfo
	regex     *regexp.Regexp
	linesCh   chan string
}

func New(logParser logParser, geoInfo geoInfo, linesCh chan string) *Report {
	return &Report{
		regex:     regexp.MustCompile(logPatter),
		logParser: logParser,
		geoInfo:   geoInfo,
		linesCh:   linesCh,
	}
}

func (r *Report) GetReport() {
	for line := range r.linesCh {
		lineLog := r.logParser.Parse(line)
		if !r.shouldExclude(lineLog.RequestPath) {
			record := r.geoInfo.GetIPInfo(lineLog.IP)
			log.Println(record)
		}
	}
}

func (r *Report) shouldExclude(requestPath string) bool {
	return r.regex.MatchString(requestPath)
}

/*
Ignore all requests for images, CSS, and JavaScript. This is any request path beginning with:

/[a-f0-9]+/css/
/[a-f0-9]+/images/
/[a-f0-9]+/js/
/entry-images/
/images/
/user-images/
/static/
/robots.txt
/favicon.ico
Please note that the first three path prefixes are regular expressions. The remaining are plain text.

Also ignore all paths ending in .rss or .atom.
*/
