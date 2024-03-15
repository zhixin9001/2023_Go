// create hello world
package main

// func Max(n int, m int) int {
// 	if n > m {
// 		return n
// 	} else {
// 		return m
// 	}
// }

// func main() {
// 	fmt.Printf("%d\n", Max(5, 6))
// }

// func Max(n int, m int, fn func(int)) {
// 	if n > m {
// 		fn(n)
// 	} else {
// 		fn(m)
// 	}
// }

// func main() {
// 	Max(5, 6, func(x int) {
// 		println(x)
// 	})
// }

// 拦截panic

// func bar() {
// 	println("raise a panic")
// 	panic(-1)
// }

// func foo() {
// 	defer func() {
// 		if e := recover(); e != nil {
// 			println("recovered from a panic")
// 		}
// 	}()
// 	bar()
// }

// func main() {
// 	foo()
// 	println("main exit normally")
// }

// // 修改函数返回值
// func foo(a int, b int) (x int, y int) {
// 	defer func() {
// 		x = x * 2
// 		y = y * 3
// 	}()
// 	x = a + 5
// 	y = b + 6
// 	return
// }

// func main() {
// 	println(foo(1, 2))
// }
