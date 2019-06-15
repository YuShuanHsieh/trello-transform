package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const outputDepth = 3

func GetLogger(name string) Logger {
	if name != "" {
		name = "[" + name + "] "
	}
	ret := log.New(os.Stdout, name, log.Lshortfile|log.LstdFlags)
	return SimpleLoggerWrapper(ret)
}

func sformat(pre string, v ...interface{}) string {
	sb := strings.Builder{}
	sb.WriteString(pre)
	sb.WriteRune(' ')
	sb.WriteString(fmt.Sprint(v...))
	return sb.String()
}

func sformatf(pre, f string, v ...interface{}) string {
	sb := strings.Builder{}
	sb.WriteString(pre)
	sb.WriteRune(' ')
	sb.WriteString(fmt.Sprintf(f, v...))
	return sb.String()
}

// SimpleLogger should be implimentable by log.Logger, and most standard loggers
type SimpleLogger interface {
	Output(calldepth int, s string) error // Is defined in log.Logger
}

// The common functions in many log freameworks, so we can start using them
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warning(v ...interface{})
	Warningf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

// Ref to standard logger, as our logger
type simpleLoggerWrapper struct {
	log SimpleLogger
}

func SimpleLoggerWrapper(sl SimpleLogger) Logger {
	return &simpleLoggerWrapper{log: sl}
}

func (l simpleLoggerWrapper) Print(v ...interface{}) {
	l.log.Output(outputDepth, fmt.Sprint(v...))
}

func (l simpleLoggerWrapper) Printf(format string, v ...interface{}) {
	l.log.Output(outputDepth, fmt.Sprintf(format, v...))
}

func (l simpleLoggerWrapper) Println(v ...interface{}) {
	l.log.Output(outputDepth, fmt.Sprintln(v...))
}

func (l simpleLoggerWrapper) Debug(v ...interface{}) {
	l.log.Output(outputDepth, sformat("[DEBUG] ", v...))
}

func (l simpleLoggerWrapper) Debugf(f string, v ...interface{}) {
	l.log.Output(outputDepth, sformatf("[DEBUG] ", f, v...))
}

func (l simpleLoggerWrapper) Info(v ...interface{}) {
	l.log.Output(outputDepth, sformat("[INFO] ", v...))
}

func (l simpleLoggerWrapper) Infof(f string, v ...interface{}) {
	l.log.Output(outputDepth, sformatf("[INFO] ", f, v...))
}

func (l simpleLoggerWrapper) Warning(v ...interface{}) {
	l.log.Output(outputDepth, sformat("[WARNING] ", v...))
}

func (l simpleLoggerWrapper) Warningf(f string, v ...interface{}) {
	l.log.Output(outputDepth, sformatf("[WARNING] ", f, v...))
}

func (l simpleLoggerWrapper) Error(v ...interface{}) {
	l.log.Output(outputDepth, sformat("[ERROR] ", v...))
}

func (l simpleLoggerWrapper) Errorf(f string, v ...interface{}) {
	l.log.Output(outputDepth, sformatf("[ERROR] ", f, v...))
}
