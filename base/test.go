package base

import "fmt"

func test() {
	// 1. 常量
	const (
		a = 1
		b = 3.14
	)
	fmt.Println(a, b)

	const (
		c int     = 10
		d float64 = 3.14
	)
	fmt.Println(c, d)

	var (
		e int
		f float64
	)
	e = 10
	f = 3.1415
	fmt.Println(e, f)

	g := 12
	h := 12.12
	fmt.Println(g, h)

	// iota
	// 1. iota常量自动生成器,每个一行，自动加一
	// 2. iota给常量赋值使用
	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)
	fmt.Println(a1, b1, c1)

	// iota遇到const会重置为0
	const d1 = iota
	const f1 = iota
	fmt.Println(d1, f1)

	// 在一个括号里可以只写一个iota
	const (
		a2 = iota
		b2
		c2
	)
	fmt.Println(a2, b2, c2)

	// iota在同一行值一样
	const (
		a3         = iota
		j1, j2, j3 = iota, iota, iota
		e3         = iota
	)
	fmt.Println(a3)
	fmt.Println(j1, j2, j3)
	fmt.Println(e3)

	// i1, i2, i3 := 1, 2.1, 3

}
