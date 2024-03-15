package main

import "fmt"

func main() {
	// 	fmt.Println("hello world")

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
	// s := make([]byte, 5)
	// fmt.Println(s) //自动创建对应的底层数组

	// u := [10]byte{1, 2, 3, 4}
	// us := u[1:9] //[low: high], high不包含
	// fmt.Println(us)
	// us[0] = 22

	// fmt.Println(u)
	// fmt.Println(us) //引用，一起改变
	// uss := us[1:3]  //再切片 reslicing
	// fmt.Println(uss)

	// var s []int
	// fmt.Println(len(s), cap(s)) //0 0
	// s = append(s, 11)
	// fmt.Println(len(s), cap(s)) //1 1
	// s = append(s, 12)
	// fmt.Println(len(s), cap(s)) //2 2
	// s = append(s, 12)
	// fmt.Println(len(s), cap(s)) //3 4 ，动态扩容，分配新数组，绑定底层数组会更换
	// s = append(s, 14)
	// fmt.Println(len(s), cap(s))
	// s = append(s, 15)
	// fmt.Println(len(s), cap(s))
	// //动态扩容，也会导致性能问题，建议再分配切片时指定容量
	// s1 := make([]int, 1, 10)
	// fmt.Println(len(s1), cap(s1))

	// Map
	// 无序键值对，value没有限制，key必须时实现==和!=的类型，
	// var m map[string]int
	// m["key"] = 1 //map不是零值可用的,panic: assignment to entry in nil map
	//Map声明、赋值
	// var m1 = map[int]int{
	// 	1: 1,
	// 	2: 2,
	// 	3: 3,
	// }
	// fmt.Println(m1)
	// var m2 = make(map[int]int)
	// m2[1] = 1
	// m2[2] = 2
	// fmt.Println(m2)
	// // len
	// fmt.Println(len(m1), len(m2))
	// // get value
	// fmt.Println(m2[1])
	// // delete
	// delete(m2, 1)
	// fmt.Println(m2)

	// delete(m2, 3) // 3不存在，但不会报错
	// fmt.Println(m2)
	//遍历
	// for k, v := range m1 {
	// 	fmt.Println(k, v) //遍历次序不能保证
	// }

	// 	// string
	// 	//不可变
	// 	//零值可用
	// 	var s string
	// 	fmt.Println("s=" + s + ",")
	// 	multileLine := `line1
	// line2
	// line3`
	// 	fmt.Println(multileLine)

	// //string builder
	// var b strings.Builder
	// b.WriteString("line1")
	// b.WriteString("line2")
	// fmt.Println(b.String())

	// break只跳出当前层级
	// exit := make(chan interface{})

	// go func() {
	// 	for {
	// 		select {
	// 		case <-time.After(time.Second):
	// 			fmt.Println("tick")
	// 		case <-exit:
	// 			fmt.Println("exiting...") //只跳出select,而不是for
	// 			break
	// 		}
	// 	}
	// 	fmt.Println("exit!")
	// }()

	// time.Sleep(3 * time.Second)
	// exit <- struct{}{}

	// // wait child goroutine exit
	// time.Sleep(3 * time.Second)

	// // break需要配合loop使用
	// exit := make(chan interface{})

	// go func() {
	// loop:
	// 	for {
	// 		select {
	// 		case <-time.After(time.Second):
	// 			fmt.Println("tick")
	// 		case <-exit:
	// 			fmt.Println("exiting...")
	// 			break loop
	// 		}
	// 	}
	// 	fmt.Println("exit!")
	// }()

	// time.Sleep(3 * time.Second)
	// exit <- struct{}{}

	// // 等待子goroutine退出
	// time.Sleep(3 * time.Second)

	// fallthrough与case表达式列表
	n := 1
	// switch n {
	// case 1:
	// 	fallthrough
	// case 3:
	// 	fallthrough
	// case 5:
	// 	fallthrough
	// case 7:
	// 	fmt.Println("odd")
	// case 2:
	// 	fallthrough
	// case 4:
	// 	fallthrough
	// case 6:
	// 	fallthrough
	// case 8:
	// 	fmt.Println("even")
	// }

	//更推荐的写法
	switch n {
	case 1, 3, 5, 7:
		fmt.Println("odd")
	case 2, 4, 6, 8:
		fmt.Println("even")
	}
}
