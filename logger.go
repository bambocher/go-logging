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
	"os"
	"sync"
)

var loggers = make(map[string]*Logger)

type Logger struct {
	mutex    sync.Mutex
	name     string
	level    int
	handlers []Handler
}

func (logger *Logger) Log(level int, args ...interface{}) error {

	if logger.level < level {
		return nil
	}

	message := fmt.Sprint(args[0])

	if len(args) != 0 {
		message = fmt.Sprintf(message, args[1:]...)
	}

	record := NewRecord(level, logger.name, message)

	for handler := range logger.handlers {
		if logger.handlers[handler].GetLevel().min <= record.levelNo && record.levelNo <= logger.handlers[handler].GetLevel().max {
			logger.handlers[handler].Handle(record)
		}
	}

	return nil
}

func (logger *Logger) SetName(name string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	if _, ok := loggers[logger.name]; ok {
		delete(loggers, logger.name)
	}

	logger.name = name

	loggers[logger.name] = logger
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
		return errors.New(fmt.Sprintf("Unknown level type %v. Expected a string or int.", level))
	}

	return nil
}

func (logger *Logger) GetLevel() int {
	return logger.level
}

func (logger *Logger) SetHandlers(handlers []Handler) error {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	logger.handlers = handlers

	return nil
}

func (logger *Logger) GetHandlers() []Handler {
	return logger.handlers
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

	stdout := GetStreamHandler("stdout", os.Stdout)
	stdout.SetLevel(WARNING, NOTSET)

	stderr := GetStreamHandler("stderr", os.Stderr)
	stderr.SetLevel(PANIC, ERROR)

	loggers[name] = &Logger{
		name:     name,
		level:    NOTSET,
		handlers: []Handler{stdout, stderr},
	}

	return loggers[name]
}
