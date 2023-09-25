package main

import (
	"fmt"

	"github.com/jquirke/linkerdeadcodeexamples/libdemo1"
)

type barface interface {
	Bar() string
}

func main() {
	var barface barface = libdemo1.NewFoo("bla")

	fmt.Println("Bar() returns + " + barface.Bar())
}
