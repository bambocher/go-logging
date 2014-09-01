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

var loggers = make(map[string]*Logger)

var root = GetLogger("root")

func GetLogger(name string) *Logger {
	if logger, ok := loggers[name]; ok {
		return logger
	}

	logger := &Logger{
		name:    name,
		level:   DEBUG,
		format:  "[{date}][{level}][{file}:{line}] {message}\n",
		datefmt: "2006-01-02 15:04:05", // http://golang.org/pkg/time/#pkg-constants
	}

	loggers[name] = logger
	return logger
}

func GetName() string {
	return root.GetName()
}

func SetLevel(level int) {
	root.SetLevel(level)
}

func GetLevel() int {
	return root.GetLevel()
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
