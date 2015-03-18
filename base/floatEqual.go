package main

import (
	"fmt"
)

func main() {
	str := "Hello, 中国"
	for i, ch := range str {
		//ch是rune类型:
		fmt.Println(i, ch)
	}
}
