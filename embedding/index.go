package embedding

import "fmt"

type Base struct {
	 name string
}

type Entity struct {
	Base
}

func (b *Base) play() {
	fmt.Println("base")
}

func Test() {
	e := &Entity{}
	// https://hackthology.com/object-oriented-inheritance-in-go.html
	e.play() // will call Base's play without downwards-shadowing.
}
