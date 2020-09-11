/**
 * @Auth: treefei
 * @Date: 2020/9/10
 * @Time: 7:41 下午
 */

package main

func main() {

}

//更好理解代码返回参数
func returnParams1() (int, string) {
	i := 1
	y := "test"
	return i, y
}

//消除冗余代码
func returnParams2() (i int, y string) {
	i = 1
	y = "test"
	return
}

//有这种写法，但是没什么必要这么写
func returnParams3() (i int, y string) {
	i = 1
	y = "test"
	return i, y
}
