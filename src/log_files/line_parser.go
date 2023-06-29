package log_files

import (
	"strings"
	"time"
)

const iso8601Layout = "2006-01-02T15:04:05Z"

type LogLine interface {
	Timestamp() (int64, error)
	Raw() string
}

type LogLineImpl struct {
	line string
}

func NewLogLineImpl(line string) *LogLineImpl {
	return &LogLineImpl{line: line}
}

func (l *LogLineImpl) Timestamp() (int64, error) {
	logLineParts := strings.Split(l.line, ",")
	datetime := logLineParts[0]
	return l.parseDatetime(datetime)
}

func (l *LogLineImpl) Raw() string {
	return l.line
}

func (l *LogLineImpl) parseDatetime(datetime string) (int64, error) {
	dateTime, err := time.Parse(iso8601Layout, datetime)
	if err != nil {
		return 0, err
	}
	return dateTime.Unix(), nil
}
