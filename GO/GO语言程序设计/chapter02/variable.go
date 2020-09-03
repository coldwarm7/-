/**
 * @Auth: treefei
 * @Date: 2020/8/31
 * @Time: 3:29 下午
 * @Desc: 关于变量的重复声明
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

const (
	openFilePath   = "~/demo1"
	createFilePath = "~/demo2"
)

func main() {
	//variableDemo1()
	//variableDemo3()
	err := variableDemo4()
	if err != nil {
		fmt.Println("err: ", err)
	}
}

func variableDemo1() {
	openFile, err := os.Open(openFilePath) //这里声明了两个变量openFile,以及err
	if err != nil {
		return
	}
	defer openFile.Close()

	createFile, err := os.Create(createFilePath) //这里声明了一个变量，createFile,随后给之前已经声明过的err赋值
	if err != nil {
		return
	}
	defer createFile.Close()
}

//变量的覆盖与作用域
func variableDemo3() {
	x := 1
	println(x) // 1
	{
		println(x) // 1
		x := 2     //这里定义了新的变量x
		println(x) // 2	// 新的 x 变量的作用域只在代码块内部
	}
	println(x) // 1
}

//下面两个demo无意义，忽略
func variableDemo2() error {
	var wg sync.WaitGroup
	openFile, err := os.Open(openFilePath)
	if err != nil {
		return err
	}
	defer openFile.Close()

	wg.Add(1)
	go func() {
		defer wg.Done()
		name, err := ReturnErr()
		if err != nil {
			return
		}
		fmt.Println(name)
	}()
	wg.Wait()

	return nil

}

func variableDemo4() error {
	name1, err := ReturnNoErr()
	if err != nil {
		return err
	}
	{
		name2, err := ReturnErr()
		if err != nil {
			return err //在这里err就会直接返回
		}
		fmt.Println(name2)
	}
	fmt.Println(name1)
	return nil
}

//还有一种情况，err的重复定义会导致程序报错，一时想不起来了，待补充

func ReturnErr() (name string, err error) {
	name = "name"
	err = errors.New("err")
	return
}

func ReturnNoErr() (string, error) {
	return "name1", nil
}
