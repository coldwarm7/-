/**
 * @Auth: treefei
 * @Date: 2020/9/4
 * @Time: 10:43 上午
 * @Desc: 类型声明
 * @参考 :http://c.biancheng.net/view/25.html
 */

package main

import "fmt"

func main() {
	typeDemo()
}

//intDemo1,intDemo2是两种新类型，具有int类型的特性
type intDemo1 int
type intDemo2 int

//给int取一个别名intAlias，本质还是int类型
type intAlias = int

func typeDemo() {
	var i intDemo1
	var y intDemo2

	i = 1
	y = 1
	//if i == y {  //虽然i和y底层都是int类型，但是一个是类型intDemo1,一个是类型intDemo2,类型不同，不能直接比较
	//
	//}
	fmt.Printf("a type: %T\n", i) //a type: main.intDemo1
	fmt.Printf("a type: %T\n", y) //a type: main.intDemo2

	if i == intDemo1(y) { //可以通过显式转换将y的intDemo2类型转为intDemo1类型，再进行比较
		fmt.Println(i)
	}
}

func typeAliasDemo() {
	var i intAlias
	var y intAlias
	i = 1
	y = 1
	if i == y { //因为intAlias都是int的别名，本质还是int，所以可以直接比较
		fmt.Println("success")
	}
}
