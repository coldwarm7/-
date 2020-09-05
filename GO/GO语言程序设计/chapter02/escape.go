/**
 * @Auth: treefei
 * @Date: 2020/9/4
 * @Time: 10:35 上午
 * @Desc: 变量逃逸。每一次变量逃逸都需要一次额外的内存分配过程，不利于gc
 */

package main

import "fmt"

var global *int

func main() {
	f()
	g()
	fmt.Println(global)  //0xc00001c090
	fmt.Println(*global) //1
}

//x从f()中逃逸
func f() {
	var x int
	x = 1
	global = &x
}

//*y没有从g中逃逸
func g() {
	y := new(int)
	*y = 1
}
