package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name           string
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying....")
}

func main() {
	sparrow := &Bird{"lcore", 24}
	s := reflect.ValueOf(sparrow).Elem()
	typeofT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeofT.Field(i).Name, f.Type(), f.Interface())
	}
}
