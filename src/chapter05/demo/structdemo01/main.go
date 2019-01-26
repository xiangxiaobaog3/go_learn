package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map

}

func main() {
	var p1 Person

	// 使用切片和map 一定要make 声明一个空间
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"

	fmt.Println(p1)
}
