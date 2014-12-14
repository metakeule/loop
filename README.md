loop
====

loop io.Reader for golang

[![Build Status](https://secure.travis-ci.org/metakeule/loop.png)](http://travis-ci.org/metakeule/loop)

Usage
-----

```go
package main

import (
	"fmt"
	"gopkg.in/metakeule/loop.v2"
)

func main() {
	l := loop.New([]byte("abc..."))

	// ad infinitum abc...abc...abc...
	for {
		b := make([]byte, 2)
		l.Read(b)
		fmt.Print(string(b))
	}
}
```