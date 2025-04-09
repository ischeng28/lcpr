/*
 * @lc app=leetcode.cn id=2761 lang=golang
 * @lcpr version=30204
 *
 * [2761] 和等于目标值的质数对
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
const mx = 1e6 + 1

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

func findPrimePairs(n int) [][]int {
	if n <= 2 {
		return [][]int{}
	}
	ans := [][]int{}

	// 只需要遍历到 n/2 即可
	for i := 2; i <= n/2; i++ {
		if !notPrime[i] && !notPrime[n-i] {
			ans = append(ans, []int{i, n - i})
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
// 2\n
// @lcpr case=end

*/

