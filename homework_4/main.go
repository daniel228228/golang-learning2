package main

import (
	"errors"
	"io"
	"os"

	"homework_4/logger"
	"homework_4/writer"
)

type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Print(args ...any)
	Printf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	SetOutput(writer io.Writer)
	SetDebugMode(debugMode bool)
}

type app struct {
	log  Logger
	file *os.File
}

func main() {
	app := &app{
		logger.NewLogger(),
		openFile("output.log"),
	}
	defer app.file.Close()

	test := func() {
		app.log.Debug("debug test")
		app.log.Debugf("%s %d %+v %T", "debugf test:", 123, struct{ a, b, c int }{0, 1, 2}, app)
		app.log.Error("error test")
		app.log.Errorf("errorf test: %v", errors.New("any error"))
		app.log.Print("print test")
		app.log.Printf("printf test: %#v %t %s", struct{}{}, false, "abcdef")
	}

	app.log.SetOutput(os.Stdin)
	app.log.SetDebugMode(false)
	test()

	app.log.SetOutput(app.file)
	app.log.SetDebugMode(true)
	test()

	app.log.SetOutput(writer.CustomWriter)
	test()
}

func openFile(filename string) *os.File {
	if file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644); err != nil {
		panic(err)
	} else {
		return file
	}
}
