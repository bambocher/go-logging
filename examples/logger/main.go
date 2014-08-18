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

package main

import (
	log "github.com/bambocher/go-logging"
)

func main() {
	test_log := log.GetLogger("test")
	test_log.SetLevel(log.TRACE)
	test_log.SetFormat("[{asctime}][{levelname}][{name}] {message}\n")
	test_log.SetDateFormat("2006-01-02 15:04:05")

	test_log.Trace("Trace message.")
	test_log.Debug("Debug message.")
	test_log.Info("Informational message.")
	test_log.Notice("Notice message.")
	test_log.Warning("Warning message.")
	test_log.Error("Error message.")

	test2_log := log.GetLogger("test2")
	test2_log.SetLevel(log.NOTICE)
	test2_log.SetFormat("[{asctime}][{levelname}][{name}] {message}\n")
	test2_log.SetDateFormat("2006-01-02")

	test2_log.Trace("Trace message.")
	test2_log.Debug("Debug message.")
	test2_log.Info("Informational message.")
	test2_log.Notice("Notice message.")
	test2_log.Warning("Warning message.")
	test2_log.Error("Error message.")
}
