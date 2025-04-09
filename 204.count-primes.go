/*
 * @lc app=leetcode.cn id=204 lang=golang
 * @lcpr version=30204
 *
 * [204] 计数质数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
const mx = 5*1e6 + 1

var notPrime = [mx]bool{true, true}

func init() {
	for i := 2; i*i < mx; i++ {
		if !notPrime[i] {
			for j := i * i; j < mx; j += i {
				notPrime[j] = true
			}
		}
	}
}

func countPrimes(n int) int {
	ans := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			ans++
		}
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

