package main

import (
	"fmt"
)

//这里go语言编程书籍上有一个bug
func test(array [5]int) {
	array[0] = 10
	fmt.Println(array)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	test(array)
	fmt.Println(array)
}
