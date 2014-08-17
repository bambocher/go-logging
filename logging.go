// go-logging - Logging library for Go
//
// Copyright (c) 2014 Dmitry Prazdnichnov <dp@bambucha.org>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// RFC 5424 http://tools.ietf.org/html/rfc5424 defines eight severity levels
	NOTSET = 255
	TRACE  = 8
	// Debug-level messages. Info useful to developers for debugging the
	// application, not useful during operations.
	DEBUG = 7
	// Informational messages. Normal operational messages - may be harvested
	// for reporting, measuring throughput, etc. - no action required.
	INFORMATIONAL = 6
	INFO          = INFORMATIONAL
	// Normal but significant condition. Events that are unusual but not error
	// conditions - might be summarized in an email to developers or admins to
	// spot potential problems - no immediate action required.
	NOTICE = 5
	// Warning conditions. Warning messages, not an error, but indication that
	// an error will occur if action is not taken, e.g. file system 85% full -
	// each item must be resolved within a given time.
	WARNING = 4
	WARN    = WARNING
	// Error conditions. Non-urgent failures, these should be relayed to
	// developers or admins; each item must be resolved within a given time.
	ERROR = 3
	ERR   = ERROR
	// Critical conditions. Should be corrected immediately, but indicates
	// failure in a secondary system, an example is a loss of a backup ISP
	// connection.
	CRITICAL = 2
	CRIT     = CRITICAL
	FATAL    = CRITICAL
	// Action must be taken immediately. Should be corrected immediately,
	// therefore notify staff who can fix the problem. An example would be the
	// loss of a primary ISP connection.
	ALERT = 1
	// System is unusable. A "panic" condition usually affecting multiple
	// apps/servers/sites. At this level it would usually notify all tech staff
	// on call.
	EMERGENCY = 0
	EMERG     = EMERGENCY
	PANIC     = EMERGENCY

	DEFAULT_FORMAT = "[{asctime}][{levelname}][{filename}:{lineno}] {message}\n"
	//  "2006-01-02 15:04:05.999999999 -0700 MST"
	//  "Mon Jan 2 15:04:05 -0700 MST 2006"
	DEFAULT_DATEFMT = "2006-01-02 15:04:05"
)

type Logger struct {
	name    string
	mutex   sync.Mutex
	level   int
	format  string
	datefmt string
}

var levelNames = map[int]string{
	NOTSET:   "notset",
	TRACE:    "trace",
	DEBUG:    "debug",
	INFO:     "info",
	NOTICE:   "notice",
	WARNING:  "warn",
	ERROR:    "error",
	CRITICAL: "crit",
	ALERT:    "alert",
	PANIC:    "panic",
}

var levelNumbers = map[string]int{
	"notset":        NOTSET,
	"trace":         TRACE,
	"debug":         DEBUG,
	"informational": INFORMATIONAL,
	"info":          INFO,
	"notice":        NOTICE,
	"warning":       WARNING,
	"warn":          WARN,
	"error":         ERROR,
	"err":           ERR,
	"critical":      CRITICAL,
	"crit":          CRIT,
	"fatal":         FATAL,
	"alert":         ALERT,
	"panic":         PANIC,
	"emerg":         EMERG,
	"emergency":     EMERGENCY,
}

var loggers = make(map[string]*Logger)

var root = GetLogger("root")

func GetLogger(name string) *Logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}

	logger := &Logger{
		name:    name,
		level:   DEBUG,
		format:  DEFAULT_FORMAT,
		datefmt: DEFAULT_DATEFMT,
	}

	loggers[name] = logger
	return logger
}

func GetLevelName(level int) string {
	return levelNames[level]
}

func GetLevelNumber(level string) int {
	return levelNumbers[level]
}

func (logger *Logger) Log(level int, format string, args ...interface{}) error {

	var (
		pc       uintptr
		pathname string
		filename string
		funcName string
		lineno   int
		ok       bool
		writer   io.Writer
	)

	if pc, pathname, lineno, ok = runtime.Caller(3); !ok {
		pathname = "???"
		filename = "???"
		lineno = 0
	} else {
		pcfunc := runtime.FuncForPC(pc)
		if pcfunc != nil {
			funcName = pcfunc.Name()
		} else {
			funcName = "???"
		}
		filename = path.Base(pathname)
	}

	replace := strings.NewReplacer(
		"{name}", logger.name,
		"{levelname}", GetLevelName(level),
		"{levelno}", strconv.Itoa(level),
		"{lineno}", strconv.Itoa(lineno),
		"{asctime}", time.Now().Format(logger.datefmt),
		"{filename}", filename,
		"{pathname}", pathname,
		"{funcName}", funcName,
		"{message}", fmt.Sprintf(format, args...),
	)

	buf := []byte(replace.Replace(logger.format))

	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	if logger.level >= WARNING {
		writer = os.Stdout
	}

	if logger.level < WARNING {
		writer = os.Stderr
	}

	_, err := writer.Write(buf)

	return err
}

func (logger *Logger) SetLevel(level int) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.level = level
}

func (logger *Logger) SetFormat(format string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.format = format
}

func (logger *Logger) SetDateFormat(datefmt string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.datefmt = datefmt
}

func (logger *Logger) Trace(format string, args ...interface{}) {
	if logger.level >= TRACE {
		logger.Log(TRACE, format, args...)
	}
}

func (logger *Logger) Debug(format string, args ...interface{}) {
	if logger.level >= DEBUG {
		logger.Log(DEBUG, format, args...)
	}
}

func (logger *Logger) Informational(format string, args ...interface{}) {
	if logger.level >= INFORMATIONAL {
		logger.Log(INFORMATIONAL, format, args...)
	}
}

func (logger *Logger) Info(format string, args ...interface{}) {
	if logger.level >= INFORMATIONAL {
		logger.Log(INFORMATIONAL, format, args...)
	}
}

func (logger *Logger) Notice(format string, args ...interface{}) {
	if logger.level >= NOTICE {
		logger.Log(NOTICE, format, args...)
	}
}

func (logger *Logger) Warning(format string, args ...interface{}) {
	if logger.level >= WARNING {
		logger.Log(WARNING, format, args...)
	}
}

func (logger *Logger) Warn(format string, args ...interface{}) {
	if logger.level >= WARNING {
		logger.Log(WARNING, format, args...)
	}
}

func (logger *Logger) Error(format string, args ...interface{}) {
	if logger.level >= ERROR {
		logger.Log(ERROR, format, args...)
	}
}

func (logger *Logger) Err(format string, args ...interface{}) {
	if logger.level >= ERROR {
		logger.Log(ERROR, format, args...)
	}
}

func (logger *Logger) Critical(format string, args ...interface{}) {
	if logger.level >= CRITICAL {
		logger.Log(CRITICAL, format, args...)
		os.Exit(1)
	}
}

func (logger *Logger) Crit(format string, args ...interface{}) {
	if logger.level >= CRITICAL {
		logger.Log(CRITICAL, format, args...)
		os.Exit(1)
	}
}

func (logger *Logger) Fatal(format string, args ...interface{}) {
	if logger.level >= CRITICAL {
		logger.Log(CRITICAL, format, args...)
		os.Exit(1)
	}
}

func (logger *Logger) Alert(format string, args ...interface{}) {
	if logger.level >= ALERT {
		logger.Log(ALERT, format, args...)
		os.Exit(1)
	}
}

func (logger *Logger) Emergency(format string, args ...interface{}) {
	if logger.level >= EMERGENCY {
		s := fmt.Sprintf(format, args...)
		logger.Log(EMERGENCY, s)
		panic(s)
	}
}

func (logger *Logger) Emerg(format string, args ...interface{}) {
	if logger.level >= EMERGENCY {
		s := fmt.Sprintf(format, args...)
		logger.Log(EMERGENCY, s)
		panic(s)
	}
}

func (logger *Logger) Panic(format string, args ...interface{}) {
	if logger.level >= EMERGENCY {
		s := fmt.Sprintf(format, args...)
		logger.Log(EMERGENCY, s)
		panic(s)
	}
}

func SetLevel(level int) {
	root.SetLevel(level)
}

func SetFormat(format string) {
	root.SetFormat(format)
}

func SetDateFormat(datefmt string) {
	root.SetDateFormat(datefmt)
}

func Trace(format string, args ...interface{}) {
	root.Trace(format, args...)
}

func Debug(format string, args ...interface{}) {
	root.Debug(format, args...)
}

func Informational(format string, args ...interface{}) {
	root.Informational(format, args...)
}

func Info(format string, args ...interface{}) {
	root.Info(format, args...)
}

func Notice(format string, args ...interface{}) {
	root.Notice(format, args...)
}

func Warning(format string, args ...interface{}) {
	root.Warning(format, args...)
}

func Warn(format string, args ...interface{}) {
	root.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	root.Error(format, args...)
}

func Err(format string, args ...interface{}) {
	root.Err(format, args...)
}

// Critical is equivalent to Printf() followed by a call to os.Exit(1).
func Critical(format string, args ...interface{}) {
	root.Critical(format, args...)
}

// Crit is equivalent to Printf() followed by a call to os.Exit(1).
func Crit(format string, args ...interface{}) {
	root.Crit(format, args...)
}

// Fatal is equivalent to Printf() followed by a call to os.Exit(1).
func Fatal(format string, args ...interface{}) {
	root.Fatal(format, args...)
}

// Alert is equivalent to Printf() followed by a call to os.Exit(1).
func Alert(format string, args ...interface{}) {
	root.Alert(format, args...)
}

// Emergency is equivalent to Printf() followed by a call to panic().
func Emergency(format string, args ...interface{}) {
	root.Emergency(format, args...)
}

// Emerg is equivalent to Printf() followed by a call to panic().
func Emerg(format string, args ...interface{}) {
	root.Emerg(format, args...)
}

// Panic is equivalent to Printf() followed by a call to panic().
func Panic(format string, args ...interface{}) {
	root.Panic(format, args...)
}
