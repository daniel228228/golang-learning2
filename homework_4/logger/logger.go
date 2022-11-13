package logger

import (
	"fmt"
	"io"
)

type logger struct {
	out       io.Writer
	debugMode bool
}

func NewLogger() *logger {
	return &logger{}
}

func (l *logger) Debug(args ...any) {
	if l.debugMode {
		l.Debugf("%s", fmt.Sprint(args...))
	}
}
func (l *logger) Debugf(format string, args ...any) {
	if l.debugMode {
		fmt.Fprintf(l.out, "[DEBUG] "+format+"\n", args...)
	}
}
func (l *logger) Print(args ...any) {
	l.Printf("%s", fmt.Sprint(args...))
}
func (l *logger) Printf(format string, args ...any) {
	if l.debugMode {
		fmt.Fprintf(l.out, "[DEBUG] "+format+"\n", args...)
	} else {
		fmt.Fprintf(l.out, format+"\n", args...)
	}
}
func (l *logger) Error(args ...any) {
	l.Errorf("%s", fmt.Sprint(args...))
}
func (l *logger) Errorf(format string, args ...any) {
	if l.debugMode {
		fmt.Fprintf(l.out, "[DEBUG] [ERROR] "+format+"\n", args...)
	} else {
		fmt.Fprintf(l.out, "[ERROR] "+format+"\n", args...)
	}
}
func (l *logger) SetOutput(writer io.Writer) {
	l.out = writer
}
func (l *logger) SetDebugMode(debugMode bool) {
	l.debugMode = debugMode
}
