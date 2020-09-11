/**
 * @Auth: treefei
 * @Date: 2020/9/11
 * @Time: 10:14 上午
 */

package main

import "fmt"

func main() {
	//rangDemo()
	rangDemo2()
}

func rangDemo() {
	i := make(map[int]string)
	i[1] = "q"
	i[2] = "w"
	i[3] = "3"

	fmt.Println(i)
	for _, v := range i { //value是map i的一个变量暂存处，会随着rang不停变化，相当于i value的副本，直接更好v的值不会改变i的值
		v = "t"
		fmt.Println(v)
	}
	fmt.Println(i)

	for k, _ := range i { //改动map的值正确的应该是这样
		i[k] = "t"
	}

	fmt.Println(i)
}

type people struct {
	name string
	age  int
}

//一般的使用场景是map的value是一个struct，如果想更改struct中的参数属性，需要通过map[key] = value 来修改
func rangDemo2() {
	peopleMap := getPeopleMap()

	//错误的做法，这里并不会改变peopleMap的值
	for _, v := range peopleMap {
		v.name = "test777"
	}
	fmt.Println(peopleMap) //[{test1 21} {test2 22} {test3 23}]

	//通过slice 下标来修改
	for k, _ := range peopleMap {
		peopleMap[k].name = "test777"
	}

	fmt.Println(peopleMap) //[{test777 21} {test777 22} {test777 23}]

	//还有种写法，比较笨就是了
	for k, v := range peopleMap {
		v.age = 777
		peopleMap[k] = v
	}
	fmt.Println(peopleMap) //[{test777 777} {test777 777} {test777 777}]
}

func getPeopleMap() []people {
	peopleMap := make([]people, 0)

	peopleMap = append(peopleMap, people{
		name: "test1",
		age:  21,
	})
	peopleMap = append(peopleMap, people{
		name: "test2",
		age:  22,
	})
	peopleMap = append(peopleMap, people{
		name: "test3",
		age:  23,
	})
	return peopleMap
}
