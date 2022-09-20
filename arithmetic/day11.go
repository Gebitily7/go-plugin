package arithmetic

/**
## 问题描述
小明刚刚看完电影《第39级台阶》，离开电影院的时候，他数了数礼堂前的台阶数，恰好是39级!
站在台阶前，他突然又想着一个问题：
如果我每一步只能迈上1个或2个台阶。先迈左脚，然后左右交替，最后一步是迈右脚，也就是说一共要走偶数步。那么，上完39级台阶，那么总共有多少种不同的上法呢？
面对庞大的计算量人力计算肯定不行，请你利用计算机的优势，帮助小明寻找这个答案。

1.递归性、规律性
2.建议用时5~15min

if(n==1) return 1;//只有一级台阶，输出1步
if(n==2) return 2;//两级台阶，可以先走一步再走一步，或者直接走两步，所以两种
else return fac(n-1)+fac(n-2);//总的规律性
*/
var sum = 0

func fac(n int, step int)  {
	if n < 0 {
		return
	}

	// 最后一步是偶数
	if n == 0 && step % 2 == 0 {
		sum++
	}

	fac(n - 1, step + 1) // 当前步数奇数
	fac(n - 2, step + 1) // 当前步数偶数

}

func Day11() {
	n := 39

	fac(n, 0)

	println(sum)
}

func Fi(item int) int {
	if item == 0 || item == 1 {
		return 1
	}

	return Fi(item - 1) + Fi(item - 2)
}