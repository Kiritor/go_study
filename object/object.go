package main

import (
	"fmt"
)

//Integer
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

//interface

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	var a Integer = 10
	var b LessAdder = &a //interface assignment

}
