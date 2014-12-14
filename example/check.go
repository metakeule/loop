package main

import (
	"fmt"
	"gopkg.in/metakeule/loop.v1"
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
