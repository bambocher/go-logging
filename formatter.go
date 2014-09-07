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
	"strings"
	"sync"
	"time"
)

var DefaultFormatter = NewFormatter("[{time}][{level}][{file}:{line}] {message}", "2006-01-02 15:04:05")

type Formatter struct {
	sync.Mutex
	format     string
	dateFormat string
}

func (formatter *Formatter) SetFormat(format string) {
	formatter.Lock()
	formatter.format = format
	formatter.Unlock()
}

func (formatter *Formatter) GetFormat() string {
	return formatter.format
}

func (formatter *Formatter) SetDateFormat(dateFormat string) {
	formatter.Lock()
	formatter.dateFormat = dateFormat
	formatter.Unlock()
}

func (formatter *Formatter) GetDateFormat() string {
	return formatter.dateFormat
}

func (formatter *Formatter) Format(record *Record) string {

	replace := strings.NewReplacer(
		"{logger}", record.logger.name,
		"{level}", record.level,
		"{line}", record.line,
		"{time}", time.Now().Format(formatter.dateFormat),
		"{file}", record.file,
		"{path}", record.path,
		"{function}", record.function,
		"{message}", record.message,
	)

	replaced := replace.Replace(formatter.format)
	if !strings.HasSuffix(replaced, "\n") {
		replaced += "\n"
	}

	return replaced
}

func NewFormatter(format, dateFormat string) *Formatter {
	return &Formatter{
		format:     format,
		dateFormat: dateFormat,
	}
}
