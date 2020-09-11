/**
 * @Auth: treefei
 * @Date: 2020/9/10
 * @Time: 7:47 下午
 */

package main

import "fmt"

//函数既可以当方法参数也可以当变量赋值
func main() {
	f2 := printCount
	funcDemo(1, f2)
}

func funcDemo(n int, f func(n int)) {
	f(n)
}

func printCount(n int) {
	fmt.Println(n)
}
