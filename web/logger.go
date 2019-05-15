package web

import (
	"code.google.com/p/log4go"
	"path/filepath"
	"strings"
	"easy/trace"
	"fmt"
)

const (
	RotateHour  = "HOUR"
	RotateDaily = "DAILY"
	FileWriter  = "FILE"
)

var logger *Logger

func InitLogger() {
	logger = NewLogger()
}

type Logger struct {
	logger    log4go.Logger
	logPrefix map[string]string
}

var mapLoggerLevel = map[string]log4go.Level{
	"FINEST":   log4go.FINEST,
	"FINE":     log4go.FINE,
	"DEBUG":    log4go.DEBUG,
	"TRACE":    log4go.TRACE,
	"INFO":     log4go.INFO,
	"WARNING":  log4go.WARNING,
	"ERROR":    log4go.ERROR,
	"CRITICAL": log4go.CRITICAL,
}

func fileWriter() *log4go.FileLogWriter {
	applicationName := FrameworkConfig.Application
	filePath := FrameworkConfig.Logger.LogPath
	fileName := filepath.Join(filePath, "/", applicationName)
	writer := log4go.NewFileLogWriter(fileName, FrameworkConfig.Logger.Rotate)
	writer.SetFormat(FrameworkConfig.Logger.Format)
	if strings.ToUpper(FrameworkConfig.Logger.RotateType) == RotateHour {
		writer = writer.RotateHour(true)
	} else if strings.ToUpper(FrameworkConfig.Logger.RotateType) == RotateDaily {
		writer = writer.SetRotateDaily(true)
	}
	writer.Start()
	return writer
}

func NewLogger() *Logger {
	var writer log4go.LogWriter
	if strings.ToUpper(FrameworkConfig.Logger.Writer) == FileWriter {
		writer = fileWriter()
	}
	var logLevel log4go.Level
	if level, ok := mapLoggerLevel[strings.ToUpper(FrameworkConfig.Logger.Level)]; ok {
		logLevel = level
	}
	logger := make(log4go.Logger)
	logger.AddFilter(FileWriter, logLevel, writer)
	return &Logger{logger: logger}
}

func (l *Logger) SetRecordPrefix(t *trace.Trace) {
	var logPrefix = map[string]string{
		"ModuleName":    t.ModuleName,
		"InterfaceName": t.InterfaceName,
		"TraceID":       t.TraceID,
		"ParentID":      t.ParentID,
		"SpanID":        t.SpanID,
	}
	l.logPrefix = logPrefix
}

func (l *Logger) recordPrefix(arg0 string) string {
	var prefix []string
	for key, val := range l.logPrefix {
		prefix = append(prefix, fmt.Sprintf("%s[%s]", key, val))
	}
	prefix = append(prefix, fmt.Sprintf("Msg[%s]", arg0))
	return strings.Join(prefix, " ")
}

func (l *Logger) Info(arg0 string, arg1 ...interface{}) {
	arg0 = l.recordPrefix(arg0)
	l.logger.Info(arg0, arg1)
}

func (l *Logger) Fine(arg0 string, arg1 ...interface{}) {
	arg0 = l.recordPrefix(arg0)
	l.logger.Fine(arg0, arg1)
}

func (l *Logger) Finest(arg0 string, arg1 ...interface{}) {
	arg0 = l.recordPrefix(arg0)
	l.logger.Finest(arg0, arg1)
}

func (l *Logger) Debug(arg0 string, arg1 ...interface{}) {
	arg0 = l.recordPrefix(arg0)
	l.logger.Debug(arg0, arg1)
}

func (l *Logger) Trace(arg0 string, arg1 ...interface{}) {
	arg0 = l.recordPrefix(arg0)
	l.logger.Debug(arg0, arg1)
}

func (l *Logger) Warning(arg0 string, arg1 ...interface{}) {
	arg0 = l.recordPrefix(arg0)
	l.logger.Warn(arg0, arg1)
}

func (l *Logger) Error(arg0 string, arg1 ...interface{}) {
	arg0 = l.recordPrefix(arg0)
	l.logger.Error(arg0, arg1)
}

func (l *Logger) Critical(arg0 string, arg1 ...interface{}) {
	arg0 = l.recordPrefix(arg0)
	l.logger.Critical(arg0, arg1)
}

func (l *Logger) Close() {
	l.logger.Close()
}
