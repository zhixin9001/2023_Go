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

	// //复合字面值简化赋值
	// var a [2]int
	// a[0] = 1
	// a[1] = 2
	// var b = [2]int{1, 2}
	// fmt.Println(a, b)

	//切片
	// Go语言中传递数组是纯粹的值拷贝
	// 切片之于数组就像是文件描述符之于文件，避免性能损耗
	s := make([]byte, 5)
	fmt.Println(s) //自动创建对应的底层数组

	u := [10]byte{1, 2, 3, 4}
	us := u[1:9] //[low: high], high不包含
	fmt.Println(us)
	us[0] = 22

	fmt.Println(u)
	fmt.Println(us) //引用，一起改变
	uss := us[1:3]  //再切片 reslicing
	fmt.Println(uss)
}
