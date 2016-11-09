package main

import (
  "fmt"
)

//值拷贝
// func main() {
//   a, b := 1, 2
//   A(a, b)
//   fmt.Println(a, b)
// }
//
// func A(s ...int) {
//   s[0] = 3
//   s[1] = 4
//   fmt.Println(s)
// }

//地址拷贝
// func main() {
//   s1 := []int{1, 2, 3, 4}
//   A(s1)
//   fmt.Println(s1)
// }
//
// func A(s []int) {
//   s[0] = 5
//   s[1] = 6
//   s[2] = 7
//   s[3] = 8
//   fmt.Println(s)
// }


//值拷贝
// func main() {
//   a := 1
//   A(&a)
//   fmt.Println(a)
// }
//
// func A(a *int) {
//   *a = 2
//   fmt.Println(*a)
// }

//匿名函数
// func main() {
//   a := func() {
//     fmt.Println("Func A")
//   }
//   a()
// }

//闭包
// func main() {
//   f := closure(10)
//   fmt.Println(f(1))
//   fmt.Println(f(2))
// }
//
//
// func closure(x int) func(int) int {
//   fmt.Printf("%p\n", &x)
//   return func (y int) int {
//     fmt.Printf("%p\n", &x)
//     return x + y
//   }
// }

//defer
// func main() {
//   for i := 0; i < 3; i++ {
//     defer func ()  {
//       fmt.Println(i)
//     }()
//   }
// }

func main() {
    A()
    B()
    C()
}

func A() {
    fmt.Println("Func A")
}

func B() {
    defer func() {
        if err:=recover(); err != nil{
            fmt.Println("Recover in B")
        }
    }()
    panic("Panic in B")
}

func C() {
    fmt.Println("Func C")
}
