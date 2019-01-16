package main

import "fmt"

func main() {
	var score float64
	fmt.Println("请输入成绩")
	fmt.Scanln(&score)

	if score > 100 {
		fmt.Println("输入有误")
	} else {
		switch int(score / 60) {
		case 1:
			fmt.Println("及格")
		case 0:
			fmt.Println("不及格")
		}
	}

}
