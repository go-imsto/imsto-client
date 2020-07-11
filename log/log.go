package log

import (
	"fmt"
	syslog "log"
)

type logger struct{}

// Default 默认实例
var Default Logger

func init() {
	syslog.SetFlags(syslog.Ltime | syslog.Lshortfile)
	Default = &logger{}
}

func Set(logger Logger) {
	if logger != nil {
		Default = logger
	}
}

func Get() Logger {
	return Default
}

func (z *logger) Debugw(msg string, keysAndValues ...interface{}) {
	syslog.Output(2, fmt.Sprint("DEBUG: "+msg, keysAndValues))
}

func (z *logger) Infow(msg string, keysAndValues ...interface{}) {
	syslog.Output(2, fmt.Sprint("INFO: "+msg, keysAndValues))
}

func (z *logger) Warnw(msg string, keysAndValues ...interface{}) {
	syslog.Output(2, fmt.Sprint("WARN: "+msg, keysAndValues))
}

func (z *logger) Errorw(msg string, keysAndValues ...interface{}) {
	syslog.Output(2, fmt.Sprint("ERROR: "+msg, keysAndValues))
}

func (z *logger) Fatalw(msg string, keysAndValues ...interface{}) {
	syslog.Fatal(msg, keysAndValues)
}
