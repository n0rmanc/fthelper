package loggers

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/n0rmanc/fthelper/shared/errors"
	"github.com/n0rmanc/fthelper/shared/utils"
)

type Logger struct {
	Name   string
	writer io.Writer
}

func (l *Logger) SetWriter(writer io.Writer) {
	l.writer = writer
}

func (l *Logger) IsPrintable(lvl LoggerLevel) bool {
	return level >= lvl
}

func (l *Logger) format(lvl, format string, msg ...interface{}) string {
	// format syntax datetime: `02-01-2006 15:04:05`
	var datetime = time.Now().Format("15:04:05")

	var arr = make([]interface{}, 3)
	arr[0] = datetime
	arr[1] = l.Name
	arr[2] = strings.ToUpper(lvl)
	arr = append(arr, msg...)

	return fmt.Sprintf("%s %-20s [%-5s] | "+format, arr...)
}

func (l *Logger) log(msg interface{}) {
	l.write(l.writer, msg)
}

func (l *Logger) write(w io.Writer, msg interface{}) {
	fmt.Fprintln(w, msg)
}

func (l *Logger) ErrorExit(handler *errors.Handler) {
	l.ErrorExitN(handler, 1)
}

func (l *Logger) ErrorExitN(handler *errors.Handler, code int) {
	if handler.HasError() {
		l.Error(handler.String())
		os.Exit(code)
	}
}

func (l *Logger) Error(format string, msg ...interface{}) {
	l.ErrorKey("error", format, msg...)
}

func (l *Logger) ErrorKey(key string, format string, msg ...interface{}) {
	if l.IsPrintable(ERROR) {
		var arr = utils.CloneArray([]interface{}{key}, msg...)
		l.log(l.format("error", "%s: "+format, arr...))
	}
}

func (l *Logger) Warn(format string, msg ...interface{}) {
	if l.IsPrintable(WARN) {
		l.log(l.format("warn", format, msg...))
	}
}

func (l *Logger) Info(format string, msg ...interface{}) {
	if l.IsPrintable(INFO) {
		l.log(l.format("info", format, msg...))
	}
}

func (l *Logger) Debug(format string, msg ...interface{}) {
	if l.IsPrintable(DEBUG) {
		l.log(l.format("debug", format, msg...))
	}
}

func (l *Logger) Log(msg interface{}) {
	if l.IsPrintable(ERROR) {
		l.log(msg)
	}
}

func (l *Logger) FLog(w io.Writer, msg interface{}) {
	if l.IsPrintable(ERROR) {
		l.write(w, msg)
	}
}

func (l *Logger) Line() {
	l.Log(LINE)
}

func (l *Logger) Newline() {
	l.Log(EMPTY)
}

func (l *Logger) Table(size uint) *Table {
	t := &Table{
		size:   size,
		logger: l,
	}

	return t.Init()
}
