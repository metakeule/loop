package main

import (
	"fmt"
	"github.com/metakeule/loop"
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
