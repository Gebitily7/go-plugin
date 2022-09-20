package money

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

type DLT struct {
	front []int
	back []int
	frontNum int
	backNum int
}

type DltV struct {
	a [5]int
	b [2]int
}

type Lottery struct {
	dlt DLT
	Data []DltV
}

func NewLottery() *Lottery {
	return &Lottery{
		dlt: DLT{
			front: genIntArr(35, 1),
			back: genIntArr(12, 1),
			frontNum: 5,
			backNum: 2,
		},
		Data: nil,
	}
}

func (l *Lottery) DaLeTou (num int) {
	if num < 1 {
		return
	}

	println(l)

	for i := 0; i < num; i++ {
		l.Data = append(l.Data, l.daLeTouSeed())
	}
}

func (l *Lottery) daLeTouSeed() DltV  {
	var data DltV

	front := make([]int, 35)
	back := make([]int, 12)

	copy(front, l.dlt.front)
	copy(back, l.dlt.back)

	var index int

	// 生成前区数字
	for i := 0; i < l.dlt.frontNum; i++ {
		// 生成随机数字
		index = randInt(len(front))
		data.a[i] = front[index]

		front = append(front[:index], front[index+1:]...)
	}

	// 生成后区数字
	for i := 0; i < l.dlt.backNum; i++ {
		// 生成随机数字
		index = randInt(len(back))

		data.b[i] = back[index]

		back = append(back[:index], back[index+1:]...)
	}

	return data
}

func genIntArr (len,first int) []int {
	temp := make([]int, len)

	for i := 0; i < len; i++ {
		temp[i] = first + i
	}

	return temp
}

func randInt(max int) int {
	if max == 0 {
		return max
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))

	num, err := strconv.Atoi(result.String())
	if err == nil{
		fmt.Printf(" ")
	}

	return num
}
