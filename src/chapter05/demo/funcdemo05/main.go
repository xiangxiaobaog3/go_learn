package main

import "fmt"

type myFunType func(int, int) int

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func main() {
	//类型推导
	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2 =", res2)
}
