package main

import "fmt"

// n1 指针
func test03(n1 *int) {
	fmt.Printf("n1的地址 = %v\n", &n1)
	*n1 = *n1 + 10
	fmt.Println("test03() n1 =", *n1)
}

func main() {
	num := 20
	// 传的是地址
	test03(&num)
	fmt.Printf("num的地址 = %v\n", &num)
	fmt.Println("main() num =", num)
}
