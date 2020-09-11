/**
 * @Auth: treefei
 * @Date: 2020/9/11
 * @Time: 10:29 上午
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	deferDemo()
}

func deferDemo() {
	i := []int{1, 2, 3, 4}

	for _, v := range i {
		go func() {
			fmt.Println(v) //4 4 4 4
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("----------------------")
	for _, v := range i {
		go func() {
			vDemo := v
			fmt.Println(vDemo) //4 4 4 4
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("----------------------")
	for _, v := range i {
		go func(vDemo int) {
			fmt.Println(vDemo) //3 4 1 2
		}(v)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("----------------------")
	for _, v := range i {
		go func() {
			defer fmt.Println(v) //4 4 4 4
		}()
	}
	time.Sleep(1 * time.Second)
}
