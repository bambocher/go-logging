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
	"os"
	"path"
)

type FileHandler struct {
	BaseHandler
	file *os.File
}

func (handler *FileHandler) Handle(record *Record) error {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	buf := handler.Format(record)

	_, err := handler.file.Write(buf)

	return err
}

func GetFileHandler(name, fileName string, dirMode, fileMode os.FileMode) (Handler, error) {
	if handler, ok := handlers[name]; ok {
		return handler, nil
	}

	err := os.MkdirAll(path.Dir(fileName), dirMode)
	if err != nil {
		return nil, errors.New("Cannot create directory: " + fileName)
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, fileMode)
	if err != nil {
		return nil, errors.New("Cannot open file: " + fileName)
	}

	handlers[name] = &FileHandler{
		BaseHandler: BaseHandler{
			name:      name,
			level:     Level{NOTSET, PANIC},
			formatter: GetFormatter("default"),
		},
		file: file,
	}

	return handlers[name], nil
}
