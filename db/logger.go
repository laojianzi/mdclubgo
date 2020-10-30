package db

import (
	"github.com/laojianzi/mdclubgo/log"
)

type Logger struct{}

func (l *Logger) Printf(format string, args ...interface{}) {
	log.ShowLine(false).Debugf(format, args...)
}
