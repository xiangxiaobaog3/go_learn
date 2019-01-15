package main

import (
	a "example2/add"
	"fmt"
	"time"
)

const Man  = 1
const Female = 2


func main() {
	fmt.Println("Name=", a.Name)
	fmt.Println("age=", a.Age)

	for {
		second := time.Now().Unix()
		if (second % Female == 0) {
			fmt.Println("female")
		} else {
			fmt.Println("man")
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
