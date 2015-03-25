/**
    Go语言共享内存的方式实现多线程
*/

package main

import (
     "fmt"
	 "sync"     //提供基本的同步基元,除了Once和WaiteGroup,大部分适用于低水平线程。
	 "runtime"  //提供go运行时环境的互操作
)

var counter int = 0

func Count (lock *sync.Mutex) {
	  //sync.Mutex 互斥锁,和线程无关
	  lock.Lock()  //枷锁,如果已经锁住,则阻塞知道lock解锁之后在枷锁。
	  counter++
	  fmt.Println(counter)
	  lock.Unlock() //解锁,如果lock未枷锁会导致运行时错误
}

func main() {
	 lock :=&sync.Mutex{}
	 for i := 0; i< 10; i++ {
		 go Count(lock)    //新开一个协程
	 }
	 for {
		 lock.Lock()
		 c :=counter
		 lock.Unlock()
		 fmt.Println("放弃处理器")
		 runtime.Gosched()   //当前go程序放弃处理器,让其他go程序运行,不挂起,之后会恢复
		 fmt.Println("恢复执行")
		 if c>=10 {
			 break
		 }else {
			 fmt.Println("现在还小于等于10呢")
		 }
	 }
}
