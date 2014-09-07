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
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func LoadJSONConfig(filename string) error {
	if len(filename) <= 0 {
		return errors.New("Empty filename")
	}

	file, err := os.Open(filename)
	if err != nil {
		return errors.New(fmt.Sprintf("Can't load json config file [%s]: %v", filename, err))
	}
	defer file.Close()

	config := Config{}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return errors.New(fmt.Sprintf("Can't parse json config file [%s]: %v", filename, err))
	}

	formatters := make(map[string]*Formatter)
	handlers := make(map[string]Handler)

	for key, value := range config.Formatters {
		formatters[key] = NewFormatter(value.Format, value.DateFormat)
	}

	for key, value := range config.Handlers {
		switch value.Type {

		case "StreamHandler":
			level := &Level{LevelToInt(value.Level.Min), LevelToInt(value.Level.Max)}
			formatter, ok := formatters[value.Formatter]
			if !ok {
				return errors.New(fmt.Sprintf("Not found formatter [%s] for handler [%s]", value.Formatter, key))
				break
			}

			property, ok := value.Properties["stream"]
			if !ok {
				return errors.New(fmt.Sprintf("Not found property [stream] for handler [%s]", key))
			}

			switch property {
			case "os.Stdout":
				handlers[key] = NewStreamHandler(level, formatter, os.Stdout)
				break
			case "os.Stderr":
				handlers[key] = NewStreamHandler(level, formatter, os.Stderr)
				break
			default:
				return errors.New(fmt.Sprintf("Unknown stream [%s] for handler [%s]", property, key))
			}
			break
		case "FileHandler":
			level := &Level{LevelToInt(value.Level.Min), LevelToInt(value.Level.Max)}

			property, ok := value.Properties["filename"]
			if !ok {
				return errors.New(fmt.Sprintf("Not found property [filename] for handler [%s]", key))
			}

			handlers[key], err = NewFileHandler(level, formatters[value.Formatter], property)
			if err != nil {
				return err
			}
			break
		case "NullHandler":
			handlers[key] = NullHandler
			break
		case "StdoutHandler":
			handlers[key] = StdoutHandler
			break
		case "StderrHandler":
			handlers[key] = StderrHandler
			break
		default:
			return errors.New(fmt.Sprintf("Unknown handler type [%s]", value.Type))
		}
	}

	for key, value := range config.Loggers {
		logger := GetLogger(key)
		logger.SetLevel(LevelToInt(value.Level))
		logger.SetHandlers()
		for x := range value.Handlers {
			handler, ok := handlers[value.Handlers[x]]
			if !ok {
				return errors.New(fmt.Sprintf("Not found handler [%s] for logger [%s]", value.Handlers[x], key))
			}
			logger.AddHandlers(handler)
		}
	}

	return nil
}
