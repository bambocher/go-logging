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
	"io"
)

type StreamHandler struct {
	BaseHandler
	stream io.Writer
}

func (handler *StreamHandler) Handle(record *Record) error {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	buf := handler.Format(record)

	_, err := handler.stream.Write(buf)

	return err
}

func GetStreamHandler(name string, stream io.Writer) Handler {
	if handler, ok := handlers[name]; ok {
		return handler
	}

	handlers[name] = &StreamHandler{
		BaseHandler: BaseHandler{
			name:      name,
			level:     Level{NOTSET, PANIC},
			formatter: GetFormatter("default"),
		},
		stream: stream,
	}

	return handlers[name]
}
