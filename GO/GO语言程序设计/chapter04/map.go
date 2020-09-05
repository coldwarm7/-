/**
 * @Auth: treefei
 * @Date: 2020/9/5
 * @Time: 3:09 下午
 * @Desc: map
 */

//关于map定义
//map并非并发安全
package main

func main() {
	mapDemo()
}

func mapDemo() {
	//关于map的三种定义
	map1 := make(map[string]int)

	var map2 = map[string]int{}

	var map3 map[string]int

	//map1中使用make是定义了map1的类型并初始化，map2同样是使用var定义了类型并初始化，可以直接使用
	map1["map1"] = 1
	map2["map2"] = 2

	//map3只是定义了map3的类型，需要初始化才可以使用
	//未初始化直接使用会造成panic: assignment to entry in nil map

	map3 = map[string]int{}
	map3["map3"] = 3

	//如何定义map看个人习惯，我是比较习惯用make

}
