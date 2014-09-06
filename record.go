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
	"path"
	"runtime"
	"sync"
)

type Record struct {
	mutex      sync.Mutex
	loggerName string
	levelName  string
	levelNo    int
	lineNo     int
	fileName   string
	pathName   string
	funcName   string
	message    string
}

func NewRecord(levelNo int, loggerName, message string) *Record {

	var (
		pc       uintptr
		lineNo   int
		fileName string
		pathName string
		funcName string
		ok       bool
	)

	if pc, pathName, lineNo, ok = runtime.Caller(4); !ok {
		pathName = "???"
		fileName = "???"
		funcName = "???"
		lineNo = 0
	} else {
		pcfunc := runtime.FuncForPC(pc)
		if pcfunc != nil {
			funcName = pcfunc.Name()
		} else {
			funcName = "???"
		}
		fileName = path.Base(pathName)
	}

	record := &Record{
		loggerName: loggerName,
		levelName:  GetLevelName(levelNo),
		levelNo:    levelNo,
		lineNo:     lineNo,
		fileName:   fileName,
		pathName:   pathName,
		funcName:   funcName,
		message:    message,
	}

	return record
}
