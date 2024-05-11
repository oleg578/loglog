package loglog

import (
	"fmt"
	"log"
	"os"
)

//log flag constants
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

var (
	lstd *log.Logger
)

func Fatal(v ...interface{}) {
	_ = lstd.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	_ = lstd.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Print(v ...interface{}) {
	_ = lstd.Output(2, fmt.Sprint(v...))
}

func Printf(format string, v ...interface{}) {
	_ = lstd.Output(2, fmt.Sprintf(format, v...))
}

func GetLogger() *log.Logger {
	return lstd
}

func New(path string, prefix string, flag int) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	lstd = log.New(f, prefix, flag)
	return nil
}
