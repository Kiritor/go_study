package main

import (
    "fmt"
	"time"
)

func Add(x,y int) {
	z:=x+y
	fmt.Println(z)
}

func main(){
	for i:=0;i<10;i++ {
		//go会在一个新的协程中并发执行
		go Add(i,i)
	}
	time.Sleep(70000)  //阻塞当前线程
}
