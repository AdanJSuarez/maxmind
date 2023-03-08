package logparser

type Log struct {
	ip            string
	ts            string
	requestMethod string
	requestPath   string
	statusCode    int64
	size          int64
}

func NewLog(ip, ts, requestMethod, requestPath string, statusCode, size int64) Log {
	return Log{
		ip:            ip,
		ts:            ts,
		requestMethod: requestMethod,
		requestPath:   requestPath,
		statusCode:    statusCode,
		size:          size,
	}
}

func (l Log) IP() string {
	return l.ip
}

func (l Log) TS() string {
	return l.ts
}

func (l Log) RequestMethod() string {
	return l.requestMethod
}

func (l Log) RequestPath() string {
	return l.requestPath
}

func (l Log) StatusCode() int64 {
	return l.statusCode
}

func (l Log) Size() int64 {
	return l.size
}
