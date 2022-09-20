package base

import "fmt"

func Inout() {
	d := 10
	f := 3.14
	c := 'a'
	s := "hello world!"

	fmt.Println(d, f, c, s)

	// %d 整型
	// %f 浮点型
	// %c 字符型
	// %s 字符串型
	fmt.Printf("%d, %f, %c, %s\n", d, f, c, s)

	// 自动格式 %v
	fmt.Printf("%v, %v, %v, %v\n", d, f, c, s)

	var in int

	_, err := fmt.Scanf("%d", &in)
	if err != nil {
		return 
	}

	fmt.Println(in)
}
