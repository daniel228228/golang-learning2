package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/TwiN/go-color"
)

type out struct {
	file *os.File
}

func NewOut(file *os.File) *out {
	return &out{
		file: file,
	}
}

func (o *out) Print(args ...any) (n int, err error) {
	fmt.Fprint(os.Stdout, colorize(fmt.Sprint(args...)))
	return fmt.Fprint(o.file, args...)
}

func (o *out) Println(args ...any) (n int, err error) {
	fmt.Fprint(os.Stdout, colorize(fmt.Sprintln(args...)))
	return fmt.Fprintln(o.file, args...)
}

func (o *out) Printf(format string, args ...any) (n int, err error) {
	fmt.Fprint(os.Stdout, colorize(fmt.Sprintf(format, args...)))
	return fmt.Fprintf(o.file, format, args...)
}

func InitOutputFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Sprintf("cannot open output file (%s): %s", filename, err.Error()))
	}

	return file
}

func colorize(s string) string {
	teamMember := regexp.MustCompile(`(\[TEAM MEMBER\])`)
	delim := regexp.MustCompile(`(==============================)`)

	s = teamMember.ReplaceAllString(s, color.InRed("$1"))
	s = delim.ReplaceAllString(s, color.InYellow("$1"))

	return s
}
