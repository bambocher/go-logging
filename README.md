golog
==========

Simple logging library for Golang.

Code status
-----------

[![Build Status](https://travis-ci.org/bambocher/golog.svg?branch=master)](https://travis-ci.org/bambocher/golog)

Installation
------------

    $ go get -u github.com/bambocher/golog

Quick-start
-----------

Minimal:

```go
package main

import log "github.com/bambocher/golog"

func main() {
    log.Info(42)
}

```

or more complicated:

```go
package main

import log "github.com/bambocher/golog"

func main() {
    log.SetLevel("notset")
    log.SetFormat("[{date}][{levelName}][{fileName}:{lineNo}] {message}")
    log.SetDateFormat("2006-01-02 15:04:05")

    log.Print("Notset message.")
    log.Trace("Trace message.")
    log.Debug("Debug message.")
    log.Info("%s message.", "Informational")
    log.Notice("Notice message.")
    log.Warning("Warning message.")
    log.Error("Error message.")
    log.Critical("Critical message.")
    log.Alert("Alert message.")
    log.Panic("Panic message.")
}

```

Examples
--------

You can find a few more examples here: [examples](examples/)

Documentation
-------------

[http://godoc.org/github.com/bambocher/golog](http://godoc.org/github.com/bambocher/golog)

License
-------

[The MIT License (MIT)](LICENSE)
