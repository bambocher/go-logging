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

package main

import (
	log "github.com/bambocher/golog"
	"os"
)

func main() {
	formatter := log.GetFormatter("default")
	formatter.SetFormat("[{date}][{levelName}][{fileName}:{lineNo}] {message}")
	formatter.SetDateFormat("2006-01-02 15:04:05")

	stdout := log.GetStreamHandler("stdout", os.Stdout)
	stdout.SetLevel(log.WARNING, log.NOTSET)
	stdout.SetFormatter(formatter)

	stderr := log.GetStreamHandler("stderr", os.Stderr)
	stderr.SetLevel(log.PANIC, log.ERROR)
	stderr.SetFormatter(formatter)

	file, err := log.GetFileHandler("file", "logs/main.log", 0777, 0666)
	if err != nil {
		panic(err)
	}

	file.SetLevel(log.PANIC, log.NOTSET)
	file.SetFormatter(formatter)

	root := log.GetLogger("root")
	root.SetLevel(log.NOTSET)
	root.SetHandlers([]log.Handler{stdout, stderr, file})

	log.Print("Notset message.")
	log.Trace("Trace message.")
	log.Debug("Debug message.")
	log.Info("Informational message.")
	log.Notice("Notice message.")
	log.Warning("Warning message.")
	log.Error("Error message.")
	log.Critical("Critical message.")
	log.Alert("Alert message.")
	log.Panic("Panic message.")
}
