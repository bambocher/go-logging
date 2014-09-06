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

var loggers = make(map[string]*Logger)
var Root = GetLogger("root")

func GetLogger(name string) *Logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}

	logger := &Logger{
		name:    name,
		level:   NOTSET,
		format:  "[{date}][{level}][{file}:{line}] {message}",
		datefmt: "2006-01-02 15:04:05", // http://golang.org/pkg/time/#pkg-constants
	}

	loggers[name] = logger
	return logger
}

func GetName() string {
	return Root.GetName()
}

func SetLevel(level interface{}) error {
	err := Root.SetLevel(level)
	return err
}

func GetLevel() int {
	return Root.GetLevel()
}

func SetFormat(format string) {
	root.SetFormat(format)
}

func GetFormat() string {
	return root.GetFormat()
}

func SetDateFormat(datefmt string) {
	root.SetDateFormat(datefmt)
}

func GetDateFormat() string {
	return root.GetDateFormat()
}

func Print(args ...interface{}) {
	Root.Print(args...)
}

func Trace(args ...interface{}) {
	Root.Trace(args...)
}

func Debug(args ...interface{}) {
	Root.Debug(args...)
}

func Info(args ...interface{}) {
	Root.Info(args...)
}

func Notice(args ...interface{}) {
	Root.Notice(args...)
}

func Warning(args ...interface{}) {
	Root.Warning(args...)
}

func Error(args ...interface{}) {
	Root.Error(args...)
}

func Critical(args ...interface{}) {
	Root.Critical(args...)
}

func Alert(args ...interface{}) {
	Root.Alert(args...)
}

// Panic is equivalent to Printf() followed by a call to panic().
func Panic(args ...interface{}) {
	Root.Panic(args...)
}
