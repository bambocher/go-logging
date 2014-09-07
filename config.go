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
	"path"
)

type ConfigHandler struct {
	Type       string
	Level      struct{ Min, Max string }
	Formatter  string
	Properties map[string]string
}

type ConfigLogger struct {
	Level    string
	Handlers []string
}

type Config struct {
	Formatters map[string]struct{ Format, DateFormat string }
	Handlers   map[string]ConfigHandler
	Loggers    map[string]ConfigLogger
}

func LoadConfig(filename string) error {
	if len(filename) <= 0 {
		return errors.New("Empty filename")
	}

	ext := path.Ext(filename)
	ext = ext[1:]

	switch ext {
	case "json":
		return LoadJSONConfig(filename)
	default:
		return errors.New(fmt.Sprintf("Unknown config file type %v, only JSON are supported types", ext))
	}

	return nil
}
