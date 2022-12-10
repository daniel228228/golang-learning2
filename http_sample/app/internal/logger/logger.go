package logger

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/TwiN/go-color"

	"http_sample/internal/config"
)

type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Print(args ...any)
	Printf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

type Parameters struct {
	File         string
	WriteConsole bool
	WriteFile    bool
	ConsoleLevel string
	FileLevel    string
}

func GetConfigParams(config *config.Config) Parameters {
	return Parameters{
		File:         config.Log.File,
		WriteConsole: config.Log.WriteConsole,
		WriteFile:    config.Log.WriteFile,
		ConsoleLevel: config.Log.ConsoleLevel,
		FileLevel:    config.Log.FileLevel,
	}
}

type Level int

const (
	Debug Level = iota
	Info
	Error
	Fatal
)

type out struct {
	log, err                     *log.Logger
	level                        Level
	debugStr, errorStr, fatalStr string
}

type logger struct {
	console *out
	file    *out
}

func NewLogger(p Parameters, file *os.File) *logger {
	setOutput := func(o *out, p string, w *logWriter) {
		switch p {
		case "fatal":
			o.level = Fatal
		case "error":
			o.level = Error
		case "info":
			o.level = Info
		case "debug":
			o.level = Debug
		default:
			panic(fmt.Sprintf("unknown log level: %s", p))
		}

		logFlag := 0

		if o.level == Debug {
			o.debugStr = "[DEBUG] "
			logFlag = log.Lshortfile
		}

		o.log = log.New(os.Stdout, "", logFlag)
		o.log.SetOutput(w)

		o.err = log.New(os.Stdout, "", log.Lshortfile)
		o.err.SetOutput(w)

		o.errorStr = "[ERROR] "
		o.fatalStr = "[FATAL] "
	}

	l := &logger{}

	if p.WriteConsole {
		l.console = &out{}
		setOutput(l.console, p.ConsoleLevel, &logWriter{})
	}
	if p.WriteFile {
		l.file = &out{}
		setOutput(l.file, p.FileLevel, &logWriter{file: file})
	}

	return l
}

func (l *logger) Debug(args ...any) {
	l.write(Debug, toLog, &params{
		args: args,
	})
}

func (l *logger) Debugf(format string, args ...any) {
	l.write(Debug, toLogf, &params{
		format: format,
		args:   args,
	})
}

func (l *logger) Print(args ...any) {
	l.write(Info, toLog, &params{
		args: args,
	})
}

func (l *logger) Printf(format string, args ...any) {
	l.write(Info, toLogf, &params{
		format: format,
		args:   args,
	})
}

func (l *logger) Error(args ...any) {
	l.write(Error, toErr, &params{
		args: args,
	})
}

func (l *logger) Errorf(format string, args ...any) {
	l.write(Error, toErrf, &params{
		format: format,
		args:   args,
	})
}

func (l *logger) Fatal(args ...any) {
	l.write(Fatal, toFatal, &params{
		args: args,
	})
}

func (l *logger) Fatalf(format string, args ...any) {
	l.write(Fatal, toFatalf, &params{
		format: format,
		args:   args,
	})
}

type params struct {
	out    *out
	format string
	args   []any
}

func (l *logger) write(level Level, f func(p *params), p *params) {
	if l.file != nil && level >= l.file.level {
		p.out = l.file
		f(p)
	}
	if l.console != nil && level >= l.console.level {
		p.out = l.console
		f(p)
	}
}

func toLog(p *params) {
	p.out.log.Output(4, fmt.Sprint(append([]any{p.out.debugStr}, p.args...)...))
}

func toLogf(p *params) {
	p.out.log.Output(4, fmt.Sprintf("%s"+p.format, append([]any{p.out.debugStr}, p.args...)...))
}

func toErr(p *params) {
	p.out.err.Output(4, fmt.Sprint(append([]any{p.out.debugStr, p.out.errorStr}, p.args...)...))
}

func toErrf(p *params) {
	p.out.err.Output(4, fmt.Sprintf("%s%s"+p.format, append([]any{p.out.debugStr, p.out.errorStr}, p.args...)...))
}

func toFatal(p *params) {
	p.out.err.Output(4, fmt.Sprint(append([]any{p.out.debugStr, p.out.fatalStr}, p.args...)...))
}

func toFatalf(p *params) {
	p.out.err.Output(4, fmt.Sprintf("%s%s"+p.format, append([]any{p.out.debugStr, p.out.fatalStr}, p.args...)...))
}

func InitLogFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(fmt.Sprintf("cannot open log file (%s): %s", filename, err.Error()))
	}

	return file
}

type logWriter struct {
	file *os.File
}

func (w *logWriter) Write(p []byte) (int, error) {
	t := time.Now()
	rec := fmt.Sprintf("%s %s", fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d.%03d", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1e6), string(p))

	if w.file != nil {
		return w.file.Write([]byte(rec))
	} else {
		return fmt.Print(colorize(rec))
	}
}

func colorize(s string) string {
	time := regexp.MustCompile(`([0-9]{4}/[0-1][0-9]/[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9].[0-9]{3})`)
	debug := regexp.MustCompile(`(\[DEBUG\])`)
	err := regexp.MustCompile(`(\[ERROR\])`)
	fatal := regexp.MustCompile(`(\[FATAL\])`)

	s = time.ReplaceAllString(s, color.InPurple("$1"))
	s = debug.ReplaceAllString(s, color.InYellow("$1"))
	s = err.ReplaceAllString(s, color.InRed("$1"))
	s = fatal.ReplaceAllString(s, color.InBold(color.InRed("$1")))

	return s
}
