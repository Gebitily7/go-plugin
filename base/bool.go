package base

import (
	"fmt"
	"unsafe"
)

func Bool() {
	var b1 bool
	fmt.Println(b1)

	b1 = true
	fmt.Println(b1)

	var b11 = false
	fmt.Println(b11)

	b2 := true
	fmt.Println(b2)

	fmt.Println(false)

	i1 := 20

	fmt.Println(i1)
	fmt.Printf("%T\n", i1)
	fmt.Println(unsafe.Sizeof(i1))

	var i2 int32
	fmt.Println(i2)
	fmt.Printf("%T\n", i2)
	fmt.Println(unsafe.Sizeof(i2))

	// 浮点型 float32 单精度， float64 双精度
	var f32 float32
	var f64 float64
	f32 = 123.135796
	f64 = 123.135796
	fmt.Println(f32, f64)

	f32 = 123.1357966
	f64 = 123.1357966
	fmt.Println(f32, f64)

	f := 3.14
	fmt.Printf("%T\n", f)

	/** 字符 */
	// 1. 单引号：字符。双引号：字符串
	var c1 byte
	c1 = 'A'
	fmt.Println(c1)
	fmt.Printf("%T\n", c1)
	fmt.Printf("%c\n", c1)

	/** 字符串 */
	//
	var str1 string
	str1 = "hello world!"
	fmt.Println(str1)

	ch1 := 'a'
	str2 := "a" // 'a''\0'
	fmt.Println(ch1, str2)

	// len 计算字符串中字符的个数，不包含\0
	// 在go语言中一个汉字3个字符
	str3 := "abc"
	fmt.Println(len(str3))

	// 字符拼接  +
	str4 := "你好"
	str5 := "世界"
	str4 += str5
	fmt.Println(str4)

	var ch2 byte
	var string1 string

	ch2 = 'a'
	string1 = "abc"
	fmt.Println(ch2, string1)

	// 字符
	// 1、单引号
	// 2、字符往往就一个字符，除了\n \t 转义字符

	// 字符串
	// 1、双引号
	// 2、字符串可以有一个或多个字符组成
	// 3、字符串都是隐藏了一个结束符\0
	fmt.Println(string1[0], string1[1], string1[2])
	fmt.Printf("%c, %c, %c\n", string1[0], string1[1], string1[2])

	/** 复数 */
	var t complex64
	t = 2.1 + 3.14i
	fmt.Println(t)

	t1 := 2.1 + 3.14i
	fmt.Printf("%T", t1)

}
