# errors - Simple Errors Wrapper

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/errors)](https://pkg.go.dev/github.com/go-zoox/errors)
[![Build Status](https://github.com/go-zoox/errors/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/errors/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/errors)](https://goreportcard.com/report/github.com/go-zoox/errors)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/errors/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/errors?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/errors.svg)](https://github.com/go-zoox/errors/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/errors.svg?label=Release)](https://github.com/go-zoox/errors/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/errors
```

## Getting Started

```go
package main

import (
    stderrors "errors"

    "github.com/go-zoox/errors"
)

func main() {
    err := f3()
    
    fmt.Println(err)
}

func f1() error {
    return errors.New("string1")
}

func f2() error {
    return errors.Wrap(f1(), "custom message f2")
}

func f3() error {
    return errors.Wrapf(f2(), "custom message f3 %s", "test")
}
```

## Inspired by
* [PumpkinSeed/errors](https://github.com/PumpkinSeed/errors) - Simple and efficient error package.
* [bnkamalesh/errors](https://github.com/bnkamalesh/errors) - A drop-in replacement for Go errors, with some added sugar! Unwrap user-friendly messages, HTTP status code, easy wrapping with multiple error types.
* [pkg/errors](https://github.com/pkg/errors) - imple error handling primitives.

## License
GoZoox is released under the [MIT License](./LICENSE).