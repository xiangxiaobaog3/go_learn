package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("你好")
	}

	var str string = "hello,world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	str = "abcok"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}
}
