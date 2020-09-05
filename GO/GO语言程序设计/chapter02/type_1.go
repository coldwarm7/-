/**
 * @Auth: treefei
 * @Date: 2020/9/4
 * @Time: 11:33 上午
 * @Desc: 在结构体成员嵌入时使用别名
 */

package main

import (
	"fmt"
	"reflect"
)

// 定义商标结构
type Brand struct {
	i int
}

// 为商标结构添加Show()方法
func (t *Brand) Show() {
	fmt.Println(t.i)
}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

//上面的Vehicle其实相当于下面的Vehicle1  好像并没有什么用
type Vehicle1 struct {
	FakeBrand1 Brand
	Brand1     Brand
}

func main() {
	// 声明变量a为车辆类型
	var a Vehicle

	// 指定调用FakeBrand的Show
	a.FakeBrand.Show()
	a.Brand.Show()
	fmt.Println("----------------------")
	a.Brand.i = 3
	a.Brand.Show()
	a.FakeBrand.Show()
	fmt.Println("----------------------")
	a.FakeBrand.i = 4
	a.Brand.Show()
	a.FakeBrand.Show()
	fmt.Println("----------------------")

	fmt.Println(fmt.Sprintf("%#v", a)) //main.Vehicle{FakeBrand:main.Brand{i:4}, Brand:main.Brand{i:3}}
	fmt.Println("----------------------")
	// 取a的类型反射对象
	ta := reflect.TypeOf(a)
	// 遍历a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.
			Name())
	}
}
