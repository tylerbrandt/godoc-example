package hello

import (
	"fmt"
	"github.com/tylerbrandt/godoc-example/lib/goodbye"
)

// Sayer is a hello-sayer
type Sayer struct {
	// The message to say
	Message string
}

// SayHello says [hello.Sayer.Message]
//
// Uses [fmt.Println] to print the message
func (h Sayer) SayHello() {
	fmt.Println(h.Message)
}

// SayHello says [Sayer.SayHello] and then [goodbye.SayGoodbye].
//
// Configures [Sayer] with [hello.Sayer.Message] set to msg
func SayHello(msg string) {
	h := Sayer{Message: msg}
	h.SayHello()
	goodbye.SayGoodbye()
}
