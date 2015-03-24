/**
    Go语言共享内存的方式实现多线程
*/

package main

import (
     "fmt"
	 "sync"
	 "runtime"
)

var counter int = 0

func Count (lock *sync.Mutex) {
	  lock.Lock()
	  counter++
	  fmt.Println(counter)
	  lock.Unlock()
}

func main() {
	 lock :=&snyc.Mutex
	 for i := 0; i< 10; i++ {
		 go Count(lock)
	 }
	 for {
		 lock.Lock()
		 c :=counter
		 lock.Unlock()
		 runtime.Gosched()
		 if c>=10 {
			 break
		 }
	 }
}
