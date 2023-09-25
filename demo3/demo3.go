package main

import (
	"fmt"
	"reflect"

	"github.com/jquirke/linkerdeadcodeexamples/libdemo1"
)

type barface interface {
	Bar() string
}

func main() {
	var barface barface = libdemo1.NewFoo("bla")

	fmt.Println("Bar() returns + " + barface.Bar())
	// On go1.21 and earlier it will result in the linker including ALL
	// methods of FooType in the binary
	reflect.TypeOf(3).MethodByName("Baz")
	m := "XYZ"
	reflect.TypeOf(3).MethodByName(m)
}
