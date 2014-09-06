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
	"bytes"
	"strconv"
	"strings"
	"sync"
	"time"
)

var formatters = make(map[string]*Formatter)

type Formatter struct {
	mutex      sync.Mutex
	name       string
	format     string
	dateFormat string
}

func (formatter *Formatter) SetFormat(format string) {
	formatter.mutex.Lock()
	defer formatter.mutex.Unlock()
	formatter.format = format
}

func (formatter *Formatter) GetFormat() string {
	return formatter.format
}

func (formatter *Formatter) SetDateFormat(dateFormat string) {
	formatter.mutex.Lock()
	defer formatter.mutex.Unlock()
	formatter.dateFormat = dateFormat
}

func (formatter *Formatter) GetDateFormat() string {
	return formatter.dateFormat
}

func (formatter *Formatter) Format(record *Record) []byte {

	replace := strings.NewReplacer(
		"{loggerName}", record.loggerName,
		"{levelName}", record.levelName,
		"{levelNo}", strconv.Itoa(record.levelNo),
		"{lineNo}", strconv.Itoa(record.lineNo),
		"{date}", time.Now().Format(formatter.dateFormat),
		"{fileName}", record.fileName,
		"{pathName}", record.pathName,
		"{funcName}", record.funcName,
		"{message}", record.message,
	)

	buf := []byte(replace.Replace(formatter.format))
	if !bytes.HasSuffix(buf, []byte("\n")) {
		buf = append(buf, '\n')
	}

	return buf
}

func GetFormatter(name string) *Formatter {
	if formatter, ok := formatters[name]; ok {
		return formatter
	}

	formatters[name] = &Formatter{
		name:       name,
		format:     "[{date}][{levelName}][{fileName}:{lineNo}] {message}",
		dateFormat: "2006-01-02 15:04:05",
	}

	return formatters[name]
}
