/*
 * @lc app=leetcode.cn id=762 lang=golang
 * @lcpr version=30204
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countPrimeSetBits(left int, right int) int {
	ans := 0
	for v := left; v <= right; v++ {
		if isPrime(onesCount(uint(v))) {
			ans++
		}
	}
	return ans
}

func onesCount(x uint) int {
	num := 0
	for x > 0 {
		x &= (x - 1)
		num++
	}
	return num
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// 6\n10\n
// @lcpr case=end

// @lcpr case=start
// 10\n15\n
// @lcpr case=end

*/

