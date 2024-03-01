package main

import "fmt"

func main() {
	fmt.Println("hello world")

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

	// const (
	// 	A, B = 1, 2
	// 	C, D
	// )
	// fmt.Println(A, B, C, D)

	// iota: const声明块（包括单行声明）中每个常量所处位置在块中的偏移值（从零开始）
	// const (
	// 	A, B = iota, iota + 10
	// 	C, D
	// )
	// fmt.Println(A, B, C, D)
	// const (
	// 	_ = iota // 略过0
	// 	A
	// 	B
	// )

	// fmt.Println(A, B)

	// const (
	// 	PI   = 3.14 //非整形常量
	// 	PI_2 = 3.1414
	// )

	// type Weekday int
	// const (
	// 	Sunday = Weekday(iota) //有类型枚举常量
	// 	Monday
	// 	Tuesday
	// )
	// fmt.Println(Sunday, Monday, Tuesday)

	//复合字面值简化赋值
	var a [2]int
	a[0] = 1
	a[1] = 2
	var b = [2]int{1, 2}
	fmt.Println(a, b)
}
