package main

import (
	c "package_example/calc"
	"fmt"
)

func main()  {
	sum := c.Add(100, 300)
	sub := c.Sub(300, 100)

	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
}