package main

import (
	"fmt"
)

func main(){
	var array []int =[]int{1,2,3,4,5,6,7,8,9,10}
	slice:=array[1:8:8]
	test(slice)
	fmt.Println(array)
}

func test(array []int){
	array[0]=1000
	fmt.Println(array)
}
