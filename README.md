golog
==========

Simple logging library for Golang.

Code status
-----------

[![Build Status](https://travis-ci.org/bambocher/golog.svg?branch=master)](https://travis-ci.org/bambocher/golog)
[![GoDoc](https://godoc.org/gopkg.in/bambocher/golog.v0?status.png)](https://godoc.org/gopkg.in/bambocher/golog.v0)

Installation
------------

    $ go get -u gopkg.in/bambocher/golog.v0

Quick-start
-----------

Minimal:

```go
package main

import log "gopkg.in/bambocher/golog.v0"

func main() {
    log.Info(42)
}

```

or more complicated:

```go
package main

import log "gopkg.in/bambocher/golog.v0"

func main() {
    log.SetLevel(log.DEBUG)
    log.SetFormat("[{time}][{level}][{file}:{line}] {message}")
    log.SetDateFormat("2006-01-02 15:04:05")
    log.AddFile("main.log")

    log.Debug("Debug message.")
    log.Info("%s message.", "Informational")
    log.Notice("Notice message.")
    log.Warning("Warning message.")
    log.Error("Error message.")
    log.Critical("Critical message.")
}

```

Examples
--------

You can find a few more examples here: [examples](examples/)

License
-------

[The MIT License (MIT)](LICENSE)
