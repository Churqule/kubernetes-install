package main

import "fmt"

// 函数类型
func test(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

// 可变参函数,变参本质上就是 slice; 只能有⼀个，且必须是最后⼀个;
func test1(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

// 匿名函数可赋值给变量,做为结构字段或者在 channel ⾥传送;

func main() {
	s1 := test(func() int {
		return 100
	})

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d,%d", 10, 20)

	fmt.Println(s1, s2)

	// 可变参数函数
	fmt.Println(test1("sum: %d", 1, 2, 3))
}
