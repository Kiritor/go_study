package main

import(
   "fmt"
   "runtime"
)

type Vector []float64
var sum float64 = 0
var c=make(chan int)
var NCPU = runtime.NumCPU()
//分配给每个CPU的计算任务
func (v Vector) DoSome(i,n int,u Vector,c chan int) {
	for ;i<n;i++ {
		sum  += v[i]
	}
	c<- 1   //任务完成标识
}

func (v Vector) DoAll(u Vector) {
	c :=make(chan int,NCPU)   //接收每个CPU完成的信号量
	runtime.GOMAXPROCS(NCPU)
	for i :=0; i<NCPU;i++ {
		go v.DoSome(i*len(v)/NCPU,(i+1)*len(v)/NCPU,u,c)
	}
	//等待所有CPU的任务完成
	for i:=0;i<NCPU;i++ {
		<-c      //能够获取到数据不阻塞,该CPU任务已经完成
		fmt.Println("该CPU的任务已经完成")
	}
}
func main(){
	v:=make(Vector,100000000)
	for i:=0;i<100000000;i++ {
		v[i] = float64(i)
	}
	v.DoAll(v)
	fmt.Println(sum)
}
