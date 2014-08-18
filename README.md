go-logging
==========

Simple logging library for Golang.

Code status
-----------

[![Build Status](https://travis-ci.org/bambocher/go-logging.svg?branch=master)](https://travis-ci.org/bambocher/go-logging)

Installation
------------

    $ go get -u github.com/bambocher/go-logging

Quick-start
-----------

```go
package main

import log "github.com/bambocher/go-logging"

func main() {
    log.SetLevel(log.TRACE)
    log.SetFormat("[{asctime}][{levelname}][{filename}:{lineno}] {message}\n")
    log.SetDateFormat("2006-01-02 15:04:05")

    log.Trace("Trace message.")
    log.Debug("Debug message.")
    log.Info("Informational message.")
    log.Notice("Notice message.")
    log.Warning("Warning message.")
    log.Error("Error message.")
}
```

Examples
--------

You can find a few more examples here: [examples](examples/)

License
-------

[The MIT License (MIT)](LICENSE)
