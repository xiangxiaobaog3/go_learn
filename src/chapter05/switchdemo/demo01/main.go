package main

import "fmt"

func test(char byte) byte {
	return char + 1
}

func main() {
	// 定义一个接收字符
	//
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	/*	switch test(key)+1 {
		case 'a':
			fmt.Println("周一")
		case 'b':
			fmt.Println("周二")
		case 'c':
			fmt.Println("周三")
		default:
			fmt.Println("输入有误")
		}*/

	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	default:
		fmt.Println("输入有误")
	}

	//switch grade := 90; {
	//	case grade > 90:
	//		fmt.Println("成绩优秀..")
	//	case grade >= 70 && grade <= 90:
	//		fmt.Println("成绩优良")
	//	case grade >= 60 && grade < 70:
	//		fmt.Println("成绩及格")
	//	default:
	//		fmt.Println("不及格")
	//}
	// 不推荐写
}
