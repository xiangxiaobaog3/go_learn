package main

import (
	"fmt"
)

func main() {
	//打印1-100之间所有是9的倍数的整数的个数及总和
	//

	var max uint64 = 100
	var count uint64 = 0
	var sum uint64 = 0
	var i uint64 = 1
	for ; i <= max; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}

	fmt.Printf("count=%v sum=%v \n", count, sum)
}
