package main

import (
	"fmt"
)

func main(){

	//1、make函数创建
	//指定slice长度,这是容量默认为长度
	var slice1 = make([]int,6,5)
	fmt.Println(slice1)

	//创建一个int slice
	//长度和容量都是5
	slice :=[]int{1,2,3,4,5}
	//初始化一个100元素的slice
	slice1 :=[]int{99:1}
	//基于数组创建一个slice
	var array [10]int = [10]int{1,2}
	slice2 :=array[:4]
}
