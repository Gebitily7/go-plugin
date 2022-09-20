package arithmetic

import (
	"fmt"
)

/**
 * 小蓝要和朋友合作开发一个时间显示的网站。
 * 在服务器上，朋友已经获取了当前的时间，用一个整数表示，值为从1970年11月11日00:00:00到当前时刻经过的毫秒数。
 * 现在，小蓝要在客户端显示出这个时间。小蓝不用显示出年月日，只需要显示出时分秒即可，毫秒也不用显示，直接舍去即可。
 * 给定一个用整数表示的时间，请将这个时间对应的时分秒输出。
 */

func Day10() {
	var ut uint64
	var h uint8
	var m uint8
	var s uint8

	_, err := fmt.Scanf("%d", &ut)
	if err != nil {
		return
	}

	t := ut / 1000

	h = uint8((t / 3600) % 24)
	m = uint8((t / 60) % 60)
	s = uint8(t % 60)

	fmt.Printf("%02d:%02d:%02d", h, m, s)
}
