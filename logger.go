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
	"fmt"
	"sync"
)

var golog = struct {
	sync.RWMutex
	loggers map[string]*Logger
}{loggers: make(map[string]*Logger)}

type Logger struct {
	sync.Mutex
	name     string
	level    int
	handlers []Handler
}

func (logger *Logger) Log(level int, args ...interface{}) error {

	if logger.level > level {
		return nil
	}

	message := fmt.Sprint(args[0])

	if len(args) != 0 {
		message = fmt.Sprintf(message, args[1:]...)
	}

	record := NewRecord(level, logger, message)

	for handler := range logger.handlers {
		handlerLevel := logger.handlers[handler].GetLevel()
		if handlerLevel.Min <= level && level <= handlerLevel.Max {
			logger.handlers[handler].Handle(record)
		}
	}

	return nil
}

func (logger *Logger) SetName(name string) {
	logger.Lock()
	defer logger.Unlock()

	golog.RLock()
	_, ok := golog.loggers[logger.name]
	golog.RUnlock()
	if ok {
		golog.Lock()
		delete(golog.loggers, logger.name)
		golog.Unlock()
	}

	logger.name = name

	golog.Lock()
	golog.loggers[logger.name] = logger
	golog.Unlock()
}

func (logger *Logger) GetName() string {
	return logger.name
}

func (logger *Logger) SetLevel(level int) {
	logger.Lock()
	logger.level = level
	logger.Unlock()
}

func (logger *Logger) GetLevel() int {
	return logger.level
}

func (logger *Logger) SetHandlers(args ...Handler) {
	logger.Lock()
	logger.handlers = args
	logger.Unlock()
}

func (logger *Logger) GetHandlers() []Handler {
	return logger.handlers
}

func (logger *Logger) AddHandlers(args ...Handler) {
	logger.Lock()
	logger.handlers = append(logger.handlers, args...)
	logger.Unlock()
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

func GetLogger(name string) *Logger {
	golog.RLock()
	logger, ok := golog.loggers[name]
	golog.RUnlock()
	if ok {
		return logger
	}

	logger = &Logger{
		name:     name,
		level:    DEBUG,
		handlers: []Handler{StdoutHandler, StderrHandler},
	}

	golog.Lock()
	golog.loggers[name] = logger
	golog.Unlock()

	return logger
}
