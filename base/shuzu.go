package main

import (
	"fmt"
)

func main() {

	//声明一个长度为10的int类型的数组
	var array0 [10]int
	//声明并初始化
	var array1 [5]int = [5]int{1,2,3,4,5}
	//声明并未索引(0开始)为1和4位置指定元素初始化
	var array2 [5]int = [5]int{1:1,4:5}
	//声明并初始化
	    array3 :=[5]int{}
	//go编译器推导数组长度2
	    array4 :=[...]{1,2}
	//二维数组
	var array5 [2][2]int
	/*
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice []int = array[5:]
     slice2 :=[...]int{6} //创建数组,但是并没有申明长度
	fmt.Println(slice,cap(slice),len(slice))
	fmt.Println(slice2,cap(slice2),len(slice2))

	slice3 :=make([]int,5,10)
	fmt.Println(cap(slice3),len(slice3))
    */
    fmt.Println(array2)
}

