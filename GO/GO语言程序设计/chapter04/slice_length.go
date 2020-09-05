/**
 * @Auth: treefei
 * @Date: 2020/9/5
 * @Time: 3:15 下午
 * @Desc: 关于slice的坑
 */

package main

import "fmt"

//一般来讲判断slice是否为空时最好使用len()去判断，而不是判断是否等于nil

func main() {
	var s1 []int //创建的是个nil slice
	if s1 == nil {
		fmt.Println("s1==nil") //s1==nil
	} else {
		fmt.Println("s1!=nil")
	}

	var arr = [5]int{}

	s1 = arr[:] //创建的是个empty slice

	if s1 == nil {
		fmt.Println("s1==nil")
	} else {
		fmt.Println("s1!=nil") //s1!=nil
	}

	s2 := make([]int, 0) //创建的是个empty slice
	if s2 == nil {
		fmt.Println("s1==nil")
	} else {
		fmt.Println("s1!=nil") //s1!=nil
	}
}
