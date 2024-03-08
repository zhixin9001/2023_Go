package main

import "fmt"

func main() {

	// // 变量声明
	// var a int32
	// var s string = "hello"
	// var i = 12
	// n := 17
	// fmt.Println(a, s, i, n)

	// var a = 17
	// var b int32 = 17
	// var c = int32(17)
	// var (
	// 	d = 17
	// 	e = float32(3.14)
	// )
	// var f int32
	// var g float32
	// fmt.Println(a, b, c, d, e, f, g)

	// var a = 1
	// b := 2
	// c := int32(3)
	// fmt.Println(a, b, c)

	const ( //无类型常量
		A = 0
		B = 1
		C = 2
	)
	// fmt.Println(A, B, C)

	type myInt int
	var a = int(5)
	var b = myInt(6)
	//fmt.Println(a + b) // 不支持隐式转换
	fmt.Println(a + int(b))

	fmt.Println(a + B) // 无类型常量的好处，不需要转换

}
