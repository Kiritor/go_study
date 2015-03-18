/**
* @Author:LCore
* @Model:main
 */
package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE:calc command [arguments] ...")
	fmt.Println("\nThe commands are :\n\t add\tAddition of two values.\n\tsqrt\tsquare root of a non-negative value.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("参数输入不正确")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("参数输入错误!")
			return
		}

		ret := simplemath.Add(v1, v2)
		fmt.Println("结果为: ", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("参数输入错误!")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("参数出入错误!")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("结果为:", ret)
	default:
		Usage()
	}
}
