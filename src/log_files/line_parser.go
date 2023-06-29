package log_files

import (
	"strings"
	"time"
)

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
	dateTime, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return 0, err
	}
	return dateTime.Unix(), nil
}
