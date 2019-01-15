package main

import (
	"fmt"
)

func add(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func list(n int)  {
	for i := 0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n - i, n)
	}
}

func main() {
	var c int
	c = add(100, 200)
	fmt.Println("add(100+200)=", c)

	list(10)
}