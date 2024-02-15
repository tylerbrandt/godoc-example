package main

import (
	"github.com/tylerbrandt/godoc-example/lib/hello"
)

func main() {
	SayHello()
	SayGoodbye()
}

// SayHello says hello
//
// This does not generate any docs!
func SayHello() {
	hello.SayHello()
}
