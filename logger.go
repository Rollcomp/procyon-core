package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	Logger = &log.Logger{
		Out:       os.Stdout,
		Formatter: NewProcyonLoggerFormatter(),
	}
)

type ProcyonLoggerFormatter struct {
	log.TextFormatter
	fullApplicationId string
	applicationId     string
}

func NewProcyonLoggerFormatter() *ProcyonLoggerFormatter {
	formatter := &ProcyonLoggerFormatter{}
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.applicationId = ""
	return formatter
}

func (f *ProcyonLoggerFormatter) Format(entry *log.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 37 // gray
	case log.WarnLevel:
		levelColor = 33 // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(
		fmt.Sprintf("[%s] \x1b[%dm%-7s\x1b[0m %s : %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), f.applicationId, entry.Message)), nil
}
