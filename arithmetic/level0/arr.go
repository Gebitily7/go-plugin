package level0

/**
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。

将最终结果插入 nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

判题标准:

系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
如果所有断言都通过，那么您的题解将被 通过。

 

示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
 

提示：

0 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums 已按 升序 排列

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func RemoveDuplicates(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 1
	}

	var i int
	k := 0
	for i = 1; i < l; i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}

// MaxProfit 最大利益
func MaxProfit(prices []int) int {
	maxProfit := 0

	l := len(prices)

	if l <= 1 {
		return maxProfit
	}

	currentStart := prices[0]
	for i := 1; i < l; i++ {
		// 如果一直降
		if  prices[i] - prices[i-1] <= 0 {
			if i-2 >= 0 && prices[i-1] - prices[i-2] > 0 {
				maxProfit = prices[i-1] - currentStart + maxProfit
			}
			currentStart = prices[i]
			continue
		}

		if prices[i] - prices[i-1] > 0 && i == l - 1 {
			maxProfit = prices[i] - currentStart + maxProfit
		}
	}

	return maxProfit
}

func Rotate1(nums []int, k int)  {
	l := len(nums)

	if l <= 1 {
		return
	}

	k = k % l

	if k == 0 {
		return
	}

	for i := 0; i < k; i++ {
		temp := nums[l - 1]
		for j := l - 1; j >= 0; j-- {
			if j == 0 {
				nums[j] = temp
			} else {
				nums[j] = nums[j - 1]
			}
		}
	}
}

func Rotate2(nums []int, k int)  {
	l := len(nums)

	if l <= 1 {
		return
	}

	k = k % l

	if k == 0 {
		return
	}

	temp := make([]int, k)

	for i := l - 1; i >= 0; i-- {
		if i >= l - k {
			temp[l - 1 - i] = nums[i]
		}
		if i < k {
			nums[i] = temp[k - 1 - i]
		} else {
			nums[i] = nums[i - k]
		}
	}
}

func Rotate3(nums []int, k int)  {
	l := len(nums)

	if l <= 1 {
		return
	}

	k = k % l

	if k == 0 {
		return
	}

	temp := make([]int, l)
	copy(temp, nums)

	for i := 0; i < l; i++ {
		if i < k {
			nums[i] = temp[l - k + i]
		} else {
			nums[i] = temp[i - k]
		}
	}
}

func Rotate4(nums []int, k int)  {
	l := len(nums)

	if l <= 1 {
		return
	}

	k %= l

	if k == 0 {
		return
	}
}