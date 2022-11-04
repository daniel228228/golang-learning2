package logger

import (
	"fmt"
	"io"
)

type logger struct {
	debugLabel string
	output     io.Writer
}

func NewLogger() *logger {
	return &logger{}
}

func (l *logger) Debug(args ...any) {
	l.Debugf("%s", fmt.Sprint(args...))
}

func (l *logger) Debugf(format string, args ...any) {
	if l.debugLabel != "" {
		fmt.Fprintf(l.output, "%s"+format+"\n", append([]any{l.debugLabel}, args...)...)
	}
}

func (l *logger) Print(args ...any) {
	l.Printf("%s", fmt.Sprint(args...))
}

func (l *logger) Printf(format string, args ...any) {
	fmt.Fprintf(l.output, "%s"+format+"\n", append([]any{l.debugLabel}, args...)...)
}

func (l *logger) Error(args ...any) {
	l.Errorf("%s", fmt.Sprint(args...))
}

func (l *logger) Errorf(format string, args ...any) {
	fmt.Fprintf(l.output, "%s%s"+format+"\n", append([]any{l.debugLabel, "[ERROR] "}, args...)...)
}

func (l *logger) SetOutput(writer io.Writer) {
	l.output = writer
}

func (l *logger) SetDebugMode(debugMode bool) {
	if debugMode {
		l.debugLabel = "[DEBUG] "
	} else {
		l.debugLabel = ""
	}
}
