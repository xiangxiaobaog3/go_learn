package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式2
	p2 := Person{"mary", 20}
	fmt.Println(p2)

	//方式3
	var p3 *Person = new(Person)
	// 因为p3是一个指针，因此标准的给字段赋值方式
	p3.Name = "smith" //也可以写成p3.Name = "smith"

	//原因：go的设计者 为了程序员使用方便，底层会对p3.Name = "smith" 进行处理
	//会给 p3 加上 取值运算 (*p3).Name = "smith"
	fmt.Println(*p3)

	//方式4
	var person *Person = &Person{}
	// 因为person是一个指针，因此标准的给字段赋值方式
	person.Name = "smith" //也可以写成person.Name = "smith"

	//原因：go的设计者 为了程序员使用方便，底层会对person.Name = "smith" 进行处理
	//会给 person 加上 取值运算 (*person).Name = "smith"
	fmt.Println(*person)
}
