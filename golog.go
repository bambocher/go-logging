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

var RootLogger = GetLogger("root")

func SetName(name string) {
	RootLogger.SetName(name)
}

func GetName() string {
	return RootLogger.GetName()
}

func SetLevel(level int) {
	RootLogger.SetLevel(level)
}

func GetLevel() int {
	return RootLogger.GetLevel()
}

func SetHandlers(args ...Handler) {
	RootLogger.SetHandlers(args...)
}

func GetHandlers() []Handler {
	return RootLogger.GetHandlers()
}

func AddHandlers(args ...Handler) {
	RootLogger.AddHandlers(args...)
}

func Debug(args ...interface{}) {
	RootLogger.Debug(args...)
}

func Info(args ...interface{}) {
	RootLogger.Info(args...)
}

func Notice(args ...interface{}) {
	RootLogger.Notice(args...)
}

func Warning(args ...interface{}) {
	RootLogger.Warning(args...)
}

func Error(args ...interface{}) {
	RootLogger.Error(args...)
}

func Critical(args ...interface{}) {
	RootLogger.Critical(args...)
}

func SetFormat(format string) {
	DefaultFormatter.SetFormat(format)
}

func SetDateFormat(dateFormat string) {
	DefaultFormatter.SetDateFormat(dateFormat)
}

func AddFile(filename string) error {
	file, err := NewFileHandler(AllLevels, DefaultFormatter, filename)
	if err != nil {
		return err
	}

	AddHandlers(file)

	return nil
}
