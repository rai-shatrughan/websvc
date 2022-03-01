package middleware

import (
	"go.uber.org/zap"
)

//Logger is wrapper for zap
type Logger struct {
	*zap.Logger
	err error
}

//New returns a instance of logger
func (l *Logger) New() {
	l.Logger, l.err = zap.NewDevelopment()
	if l.err != nil {
		l.Error("Can't initialize zap logger: %v", zap.Error(l.err))
	}
	defer l.Logger.Sync()
}
