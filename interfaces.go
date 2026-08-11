package main

import (
	"fmt"
	"math"
)

//1 task

type Shape interface {
	Area() float64
}

type Circle struct{ R float64 }

var _ Shape = Circle{}

type Rectangle struct{ W, H float64 }

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

func Describe(s Shape) {
	fmt.Println(s.Area())
}

//7 task (with error)

type ReadWriteCloser interface {
	Read()
	Write()
	Close()
}

type ReadWriter struct{ s string }

func (rwc ReadWriter) Read() {
	fmt.Println(rwc.s)
}
func (rwc ReadWriter) Write() {
	fmt.Println(rwc.s)
}
