# terrors

Small but useful errors for Go codebases.

## Installation

`go get github.com/jalavosus/terrors`

## Usage

```go
package main

import (
	"errors"
	"fmt"

	"github.com/jalavosus/terrors"
)

func recoverFunc() {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			var check *terrors.NotImplementedError
			if errors.Is(err, check) {
				fmt.Println(err.Error())
			} else if errors.Is(err, terrors.ErrNotImplemented) {
				fmt.Println(err.Error())
			}
		}
	}
}

func ThisWillPanic() {
	defer recoverFunc()
	panic(terrors.ErrNotImplemented)
}

func ThisWillPanic2() {
	defer recoverFunc()
	panic(terrors.NewNotImplementedError("ThisWillPanic2 hasn't been implemented because of the dolphins"))
}

func main() {
	ThisWillPanic()
	ThisWillPanic2()
}
```