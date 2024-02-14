package main

import (
	"github.com/tylerbrandt/godoc-example/lib/hello"
)

func main() {
	SayHello()
}

// SayHello says hello
func SayHello() {
	hello.SayHello()
}
