/**
 @Author LCore
 @Description:is prime
*/

package main

import (
   "fmt"
   "time"
)

/**
  最初浅的理解
  时间复杂度:O(n)
*/
func isPrime(num int64) bool {
   if num < 2 {
       return false
   }
   for i:=int64(2);i<num;i++ {
       if num%i==0  {
          return false
       }
   }

   return true
}

/**
  改进:去掉偶数(考虑2的特殊)
  时间复杂度:o(n/2)
*/
func isPrime2(num int64) bool {
   if num < 2 {
      return false
   }
   if num == 2 {
      return true
   }
   if num%2 == 0 {
	     return false
   }
   for i:=int64(3);i<num;i+=2 {
      if num%i== 0 {
		 return false
	  }
   }
   return true
}

/**
 素数的判断依据:
 定理:如果n不是素数,则n有满足1<d<=sqrt(n)的一个因子d
 证明:如果n不是素数,则由定义n有一个因子d满足1<d<n
     如果d>sqrt(n),则n/d是满足1<n/d<=sqrt(n）的一个因子
 o(sqrt(n)/2) 速度提高o((n-sqrt(n))/2)
*/




/**
 * 剔除因子中的重复判断
   11%3!＝0 则可以确定11%(3*i)!=0
   定理:如果n不是素数,则n有满足1<d<=sqrt(n)的一个素数因子。
   证明:1、如果n不是素数,则n有满足1<d<=sqrt(n)的一个素数因子。
       2、如果d是素数,算法终止。
       3、n＝d,转至1
*/

//primes素数递增数组
func isPrime4(primes []int64,num int64) bool {
	if num<2 {
	   return false
	}
	for i:=int64(0);primes[i]*primes[i] <= num;i++ {
	   if num%primes[i] == 0 {
	      return false
	   }
	}
	return true
}

func makePrimes(primes []int64,num int64) {

}

func isPrime3(num int64) bool {
   if num<2 {
      return false
   }
   if num==2 {
      return true
   }

   for i:=int64(3);i*i<=num;i+=2 {
	  //fmt.Println(num%i)
      if num%i == 0 {
	     return false
	  }
   }
   return true
}
func main(){
    var N int64 = 100000000
	var count int64 = 0
	var sum int64 = 0
	t1 := time.Now()
    for i:=int64(1);i<=N;i++ {
      if isPrime3(int64(i)) {
		  count+=1
		  sum +=i
          fmt.Println(i)
      }
    }
	fmt.Println(time.Now().Sub(t1))
	fmt.Println(sum)
}
