package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		break 和 continue 练习
		1、求[100,999]的水仙花数
		2、求2-100的素数
	*/
	// 1、水仙花数
	for i := 100; i < 1000; i++ {
		x := i / 100
		y := i % 100 / 10 // 或 i / 10 % 10
		z := i % 10
		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}
	fmt.Println("-----------------")
	// 个：0-9
	// 十：0-9
	// 百：1-9
	for x := 1; x < 10; x++ {
		for y := 0; y < 10; y++ {
			for z := 0; z < 10; z++ {
				n := x*100 + y*10 + z
				if x*x*x+y*y*y+z*z*z == n {
					fmt.Println(n)
				}

			}
		}
	}

	// 2、打印2-100的素数(只能被1和它自己整除)
	for i := 2; i <= 100; i++ {
		flag := true
		for j:=2;j<i;j++{
			if i % j ==0{
				flag = false
				break
			}
		}
		if flag{
			fmt.Printf("%d ", i)
		}
	}
}
