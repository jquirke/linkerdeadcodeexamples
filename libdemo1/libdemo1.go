package libdemo1

import "github.com/jquirke/linkerdeadcodeexamples/libdemo2"

type Foo struct {
	Msg string
}

// if this New function gets inlined
// then main.main will not need a type
// import of Foo.
//
// Disabling the inlining is needed
// to keep this example trivial

//go:noinline
func NewFoo(msg string) *Foo {
	return &Foo{Msg: msg}
}
func (f *Foo) Bar() string {
	return f.Msg + " foo"
}

func (f *Foo) Baz() string {
	return f.Msg + " bar"
}

func (f *Foo) Barf() string {
	return libdemo2.Barf()
}

func (f *Foo) Many() string {
	return f.Msg + " Many"
}

func (f *Foo) Other() string {
	return f.Msg + " Other"
}

func (f *Foo) Methods() string {
	return f.Msg + " Methods"
}

func (f *Foo) butOneUnexportd() string {
	return f.Msg + " butOneUnexportd"
}
