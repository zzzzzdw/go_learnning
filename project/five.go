package main

import (
	"fmt"
)

/**
* Author	: zhangdongwei@bytedance.com
* Time		: 2022-05-13
 * goto语句
 * 语法： goto LABLE
 *       LABLE: 语句块
*/
func main() {
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Println(a)
		a++
	}
}
