/**
 * @Auth: treefei
 * @Date: 2020/9/10
 * @Time: 7:42 下午
 */

package main

import "fmt"

func main() {
	retry(0)
}

func retry(count int) {
	//todo
	err := err(count)
	if err != nil && count < 5 {
		fmt.Println("retry count: ", count)
		retry(count + 1)
	}
	fmt.Println("result ", count)
}

//retry count:  0
//retry count:  1
//retry count:  2
//retry count:  3
//result  4
//result  3
//result  2
//result  1
//result  0

func err(count int) error {
	if count > 3 {
		return nil
	}
	return fmt.Errorf("count less than or equal to 3")
}
