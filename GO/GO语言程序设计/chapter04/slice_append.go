/**
 * @Auth: treefei
 * @Date: 2020/9/5
 * @Time: 2:47 下午
 * @Desc: append以及slice的容量
 */

//如果slice容量够，会直接引用原slice，将新的元素添加到后面，形成一个新的slice
//如果slice容量不够，会创建一个有足够容量的slice复制过去，再添加新元素，形成新的slice

//每次数组容量扩展的时候，都会扩展一倍的容量
//同样可以看之前提到的那个文章https://www.jianshu.com/p/1b283eeefa2b
package main

import "fmt"

func main() {
	appendDemo2()
}

func appendDemo1(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendDemo2() {
	z := make([]int, 2, 4)
	z = append(z, 1)
	z = append(z, 2)
	z = append(z, 3)
	z = append(z, 4)
	z = append(z, 5)
	z[0] = 7
	fmt.Println(z)
}
