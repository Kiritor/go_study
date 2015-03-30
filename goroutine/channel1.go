package main

import(
   "fmt"
   "time"
)

func main(){
	ch:= make(chan int)
	timeout:= make(chan bool,1)
    go func(){
		time.Sleep(1e9)  //等待一秒钟
		timeout <- true
	}()
	//通过select实现超时机制
	select {
		case <-ch:
		case <-timeout:
	}
	fmt.Println("通过Select实现超时机制,避免了死锁")
}
