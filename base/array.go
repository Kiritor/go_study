package main

import (
	"fmt"
)

func main(){

	//数组遍历
	/*for i:=0;i<len(array);i++{
		fmt.Print(array[i]," ")
	}
	//range 和匿名变量
	fmt.Println();
	for _,e :=range array {
		fmt.Print(e," ")
	}
    */
	var array [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	var array1 [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	flag :=array1==array
	fmt.Println(flag)  //true
}
