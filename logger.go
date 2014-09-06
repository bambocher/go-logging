// golog - Logging library for Go
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

package golog

import (
	"errors"
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

var loggers = make(map[string]*Logger)

type Logger struct {
	name    string
	mutex   sync.Mutex
	level   int
	format  string
	datefmt string
}

func (logger *Logger) Log(level int, args ...interface{}) error {

	if logger.level < level {
		return nil
	}

	var (
		pc       uintptr
		pathname string
		filename string
		funcName string
		message  string
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

	message = fmt.Sprint(args[0])

	if len(args) != 0 {
		message = fmt.Sprintf(message, args[1:]...)
	}

	replace := strings.NewReplacer(
		"{name}", logger.name,
		"{level}", GetLevelName(level),
		"{line}", strconv.Itoa(lineno),
		"{date}", time.Now().Format(logger.datefmt),
		"{file}", filename,
		"{path}", pathname,
		"{func}", funcName,
		"{message}", message,
	)

	buf := []byte(replace.Replace(logger.format))
	if buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}

	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	if logger.level >= WARNING {
		writer = os.Stdout
	} else {
		writer = os.Stderr
	}

	_, err := writer.Write(buf)

	return err
}

func (logger *Logger) GetName() string {
	return logger.name
}

func (logger *Logger) SetLevel(level interface{}) error {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	switch level.(type) {
	case int:
		logger.level = level.(int)
	case string:
		logger.level = GetLevelNumber(level.(string))
	default:
		return errors.New("Incorrect parameter type level, a valid string or int")
	}

	return nil
}

func (logger *Logger) GetLevel() int {
	return logger.level
}

func (logger *Logger) SetFormat(format string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.format = format
}

func (logger *Logger) GetFormat() string {
	return logger.format
}

func (logger *Logger) SetDateFormat(datefmt string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.datefmt = datefmt
}

func (logger *Logger) GetDateFormat() string {
	return logger.datefmt
}

func (logger *Logger) Print(args ...interface{}) {
	logger.Log(NOTSET, args...)
}

func (logger *Logger) Trace(args ...interface{}) {
	logger.Log(TRACE, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.Log(DEBUG, args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.Log(INFO, args...)
}

func (logger *Logger) Notice(args ...interface{}) {
	logger.Log(NOTICE, args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	logger.Log(WARNING, args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.Log(ERROR, args...)
}

func (logger *Logger) Critical(args ...interface{}) {
	logger.Log(CRITICAL, args...)
}

func (logger *Logger) Alert(args ...interface{}) {
	logger.Log(ALERT, args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.Log(PANIC, args...)
}

func GetLogger(name string) *Logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}

	loggers[name] = &Logger{
		name:     name,
		level:    NOTSET,
		handlers: []Handler{stdout, stderr},
	}

	return loggers[name]
}
