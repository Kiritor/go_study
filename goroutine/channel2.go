package main

import (
    "fmt"
)

func main(){
	var sum float64 = 0
	for i:=0;i<100000000;i++ {
		sum +=float64(i)
	}
	fmt.Println(sum)
}
