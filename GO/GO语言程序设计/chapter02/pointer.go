/**
 * @Auth: treefei
 * @Date: 2020/9/3
 * @Time: 8:17 下午
 * @Desc: 指针相关
 */

//Receiver Type 为什么推荐使用指针？

//推荐在实例方法上使用指针（前提是这个类型不是一个自定义的 map、slice 等引用类型）
//当结构体较大的时候使用指针会更高效
//如果要修改结构内部的数据或状态必须使用指针
//当结构类型包含 sync.Mutex 或者同步这种字段时，必须使用指针以避免成员拷贝
//如果你不知道该不该使用指针，使用指针！

//-----------------------------------------------------------------------

//方法参数该使用什么类型？

//map、slice 等类型不需要使用指针（自带 buf）
//指针可以避免内存拷贝，结构大的时候不要使用值类型
//值类型和指针类型在方法内部都会产生一份拷贝，指向不同
//小数据类型如 bool、int 等没必要使用指针传递
//初始化一个新类型时（像 NewEngine() *Engine）使用指针
//变量的生命周期越长则使用指针，否则使用值类型
package main

import "fmt"

func main() {
	//pointerDemo1()
	//pointerDemo2()
	//mapDemo1()

	//slice有坑，具体可以看下面的两个例子以及之前写的一个文章，图略丑，见谅https://www.jianshu.com/p/1b283eeefa2b
	//sliceDemo1()
	//sliceDemo2()
}

func pointerDemo1() {
	p := new(int)
	fmt.Println(p)  //0xc00001c090
	fmt.Println(*p) //0
	*p = 2
	fmt.Println(p)  //0xc00001c090
	fmt.Println(*p) //2
}

//变量在函数中是值传递
func pointerDemo2() {
	i := 1
	numberAdd(i)
	fmt.Println(i) //1

	pointerNumberAdd(&i)
	fmt.Println(i) //2
}

func numberAdd(i int) {
	i++
}

func pointerNumberAdd(i *int) {
	*i++
}

//map,slice自带指针，不用考虑值传递的问题
func mapDemo1() {
	i := make(map[string]int, 0)
	i["test1"] = 1
	mapAddDemo(i)
	fmt.Println(i) //map[test1:1 test2:2]
}

func mapAddDemo(i map[string]int) {
	i["test2"] = 2
}

//从这里看好像slice也是值传递，但其实也不全是
func sliceDemo1() {
	i := []string{} //等同与i := make([]string, 0)
	i = append(i, "test1")
	sliceAddDemo(i)
	fmt.Println(i) //[test1]

	slicePointerAddDemo(&i)
	fmt.Println(i) //[test1 test2]

}

//从这里看slice又是指针传递了
func sliceDemo2() {
	i := []string{"test1", "test2"} //等同与i := make([]string, 0)
	sliceModifyDemo(i)
	fmt.Println(i) //[test1 test3]
}

func slicePointerAddDemo(i *[]string) {
	*i = append(*i, "test2")
}
func sliceAddDemo(i []string) {
	i = append(i, "test2")
}

func sliceModifyDemo(i []string) {
	if len(i) >= 2 {
		i[1] = "test3"
	}
}
