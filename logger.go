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

type Logger struct {
	name    string
	mutex   sync.Mutex
	level   int
	format  string
	datefmt string
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
		"{level}", GetLevelName(level),
		"{line}", strconv.Itoa(lineno),
		"{date}", time.Now().Format(logger.datefmt),
		"{file}", filename,
		"{path}", pathname,
		"{func}", funcName,
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

func (logger *Logger) GetName() string {
	return logger.name
}

func (logger *Logger) SetLevel(level int) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.level = level
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

func (logger *Logger) Print(format string, args ...interface{}) {
	if logger.level >= NOTSET {
		logger.Log(NOTSET, format, args...)
	}
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
	if logger.level >= INFO {
		logger.Log(INFO, format, args...)
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
	if logger.level >= WARN {
		logger.Log(WARN, format, args...)
	}
}

func (logger *Logger) Error(format string, args ...interface{}) {
	if logger.level >= ERROR {
		logger.Log(ERROR, format, args...)
	}
}

func (logger *Logger) Err(format string, args ...interface{}) {
	if logger.level >= ERR {
		logger.Log(ERR, format, args...)
	}
}

func (logger *Logger) Critical(format string, args ...interface{}) {
	if logger.level >= CRITICAL {
		logger.Log(CRITICAL, format, args...)
		os.Exit(1)
	}
}

func (logger *Logger) Crit(format string, args ...interface{}) {
	if logger.level >= CRIT {
		logger.Log(CRIT, format, args...)
		os.Exit(1)
	}
}

func (logger *Logger) Fatal(format string, args ...interface{}) {
	if logger.level >= FATAL {
		logger.Log(FATAL, format, args...)
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
	if logger.level >= EMERG {
		s := fmt.Sprintf(format, args...)
		logger.Log(EMERG, s)
		panic(s)
	}
}

func (logger *Logger) Panic(format string, args ...interface{}) {
	if logger.level >= PANIC {
		s := fmt.Sprintf(format, args...)
		logger.Log(PANIC, s)
		panic(s)
	}
}
