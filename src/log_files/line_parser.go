package log_files

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type LogLineImpl struct {
	line string
}

func NewLogLineImpl(line string) *LogLineImpl {
	return &LogLineImpl{line: line}
}

func (l *LogLineImpl) Timestamp() (int64, error) {
	logLineParts := strings.Split(l.line, ",")
	if len(logLineParts) != 2 {
		return 0, errors.New(fmt.Sprintf("Invalid log line: %v", l.line))
	}
	datetime := logLineParts[0]
	return l.parseDatetime(datetime)
}

func (l *LogLineImpl) Raw() string {
	return l.line
}

func (l *LogLineImpl) IsValid() bool {
	_, err := l.Timestamp()
	if err != nil {
		log.Warnf("Invalid log line: %v; err = %v", l.line, err)
		return false
	}
	return true
}

func (l *LogLineImpl) parseDatetime(datetime string) (int64, error) {
	dateTime, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return 0, err
	}
	return dateTime.Unix(), nil
}
