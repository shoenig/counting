counting
========

Count the number of bytes passing through an io.Writer

[![Go Report Card](https://goreportcard.com/badge/go.gophers.dev/pkgs/counting)](https://goreportcard.com/report/go.gophers.dev/pkgs/counting)
[![Build Status](https://travis-ci.com/shoenig/counting.svg?branch=master)](https://travis-ci.com/shoenig/counting)
[![GoDoc](https://godoc.org/go.gophers.dev/pkgs/counting?status.svg)](https://godoc.org/go.gophers.dev/pkgs/counting)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/counting.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/shoenig/counting.svg)](LICENSE)

# Project Overview

Module `go.gophers.dev/pkgs/counting` provides a package with a utility function
for counting the number of bytes that pass through an `io.Writer`.

# Getting Started

The `counting` package can be installed by running

```bash
go get go.gophers.dev/pkgs/counting
```

#### Example Usage
```golang
f, _ := ioutil.TempFile("", "counting-")
counter := counting.NewCountingWriter(f)
compressor, _ := gzip.NewWriterLevel(counter, 9)
s := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa") // 42 bytes
u, _ := compressor.Write(s)
compressor.Close()
c := counter.Written()
fmt.Println("uncompressed bytes written", u)
fmt.Println("compressed bytes written", c)
```

# Contributing

The `go.gophers.dev/pkgs/counting` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `go.gophers.dev/pkgs/counting` module is open source under the [BSD-3-Clause](LICENSE) license.
