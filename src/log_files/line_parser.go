package log_files

import (
	"strings"
	"time"
)

const iso8601Layout = "2006-01-02T15:04:05Z"

type LogLine interface {
	timestamp() (int64, error)
	raw() string
}

type logLineImpl struct {
	line string
}

func (l *logLineImpl) Timestamp() (int64, error) {
	logLineParts := strings.Split(l.line, ",")
	datetime := logLineParts[0]
	return l.parseDatetime(datetime)
}

func (l *logLineImpl) Raw() string {
	return l.line
}

func (l *logLineImpl) parseDatetime(datetime string) (int64, error) {
	dateTime, err := time.Parse(iso8601Layout, datetime)
	if err != nil {
		return 0, err
	}
	return dateTime.Unix(), nil
}
