package arithmetic

import (
	"fmt"
	"strconv"
)

/**
 * 古代中国使用天干地支来记录当前的年份。
 * 天干一共有十个，分别为：甲（jiǎ）、乙（yǐ）、丙（bǐng）、丁（dīng）、戊（wù）、己（jǐ）、庚（gēng）、辛（xīn）、壬（rén）、癸（guǐ）。
 * 地支一共有十二个，分别为：子（zǐ）、丑（chǒu）、寅（yín）、卯（mǎo）、辰（chén）、巳（sì）、午（wǔ）、未（wèi）、申（shēn）、酉（yǒu）、戌（xū）、 亥（hài）。
 * 将天干和地支连起来，就组成了一个天干地支的年份，例如：甲子。
 * 每过一年，天干和地支都会移动到下一个。例如2021年是辛丑年。
 * 每过60年，天干会循环6轮，地支会循环5轮，所以天干地支纪年每60年轮回一次。例如1900年，1960年，2020年都是庚子年。
 * 给定一个公元纪年的年份，请输出这一年的天干地支年份。
 * 作者：知心宝贝
 * 链接：https://juejin.cn/post/7057417962197631013
 * 来源：稀土掘金
 * 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func Day08() {
	fmt.Println("请输入年份")
	year := 0
	_, err := fmt.Scanf("%d", &year)
	if err != nil {
		return
	}
	
	t := [10]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	d := [12]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}


	// 1. 2020 庚子
	//cm := 60 // 公倍数
	//x := 2020
	//
	////y := x % cm
	//fmt.Printf("%d", x % 10) // 0 => "庚"
	//fmt.Printf("%d", x % 12) // 4 => "子" -> 0 => "申"

	tOffset := 6 // 天干偏移量
	dOffset := 4 // 地支偏移量

	// 求出天干
	tIndex := year % 10 + tOffset

	if tIndex >= 10 {
		tIndex = tIndex - 10
	}

	//fmt.Printf("t%d", tIndex)

	dIndex := year % 12 - dOffset

	if dIndex < 0 {
		dIndex = dIndex + 12
	}
	//fmt.Printf("d%d", dIndex)


	fmt.Printf("%s", strconv.Itoa(year) + " => " + t[tIndex] + d[dIndex])
}
