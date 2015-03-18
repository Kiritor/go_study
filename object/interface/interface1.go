package main

import (
    "fmt"
	"one"
	"two"
)

type File struct {
}

func (f *File) Read(buf []byte) (n int, err error)
func (f *File) Write(buf []byte) (n int, err error)
func main() {
	var file1 two.IStream = new(File)
	var file2 one.ReadWriter = file1
	var file3 two.IStream = file2
	fmt.Println(file3.IStream(new []byte))
}
