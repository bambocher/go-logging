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
	"sync"
)

var handlers = make(map[string]Handler)

type Handler interface {
	SetName(name string)
	GetName() string
	SetLevel(min, max interface{}) error
	GetLevel() Level
	SetFormatter(formater *Formatter)
	GetFormatter() *Formatter
	Format(record *Record) []byte
	Handle(record *Record) error
}

// BaseHandler struct dispatch logging events to specific destinations.
type BaseHandler struct {
	mutex     sync.Mutex
	name      string
	level     Level
	formatter *Formatter
}

func (handler *BaseHandler) SetName(name string) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	if _, ok := handlers[handler.name]; ok {
		delete(handlers, handler.name)
	}

	handler.name = name

	handlers[handler.name] = handler
}

func (handler *BaseHandler) GetName() string {
	return handler.name
}

func (handler *BaseHandler) SetLevel(min, max interface{}) error {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	switch min.(type) {
	case int:
		handler.level.min = min.(int)
	case string:
		handler.level.min = GetLevelNumber(min.(string))
	default:
		return errors.New(fmt.Sprintf("Unknown level type %v. Expected a string or int.", min))
	}

	switch max.(type) {
	case int:
		handler.level.max = max.(int)
	case string:
		handler.level.max = GetLevelNumber(max.(string))
	default:
		return errors.New(fmt.Sprintf("Unknown level type %v. Expected a string or int.", max))
	}

	return nil
}

func (handler *BaseHandler) GetLevel() Level {
	return handler.level
}

func (handler *BaseHandler) SetFormatter(formater *Formatter) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()
	handler.formatter = formater
}

func (handler *BaseHandler) GetFormatter() *Formatter {
	return handler.formatter
}

func (handler *BaseHandler) Format(record *Record) []byte {
	return handler.formatter.Format(record)
}

func (handler *BaseHandler) Handle(record *Record) error {
	return nil
}
