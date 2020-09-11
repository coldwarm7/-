/**
 * @Auth: treefei
 * @Date: 2020/9/10
 * @Time: 7:52 下午
 */

package main

import "fmt"

//函数变量无法进行比较

func main() {
	f := squares()
	fmt.Println(f()) //1
	fmt.Println(f()) //4
	fmt.Println(f()) //9
	fmt.Println(f()) //16

	squares2()
	squares2()
	squares2()
	squares2()
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func squares2() {
	var x int
	x++
	fmt.Println(x)
}
