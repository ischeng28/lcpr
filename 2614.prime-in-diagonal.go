/*
 * @lc app=leetcode.cn id=2614 lang=golang
 * @lcpr version=30204
 *
 * [2614] 对角线上的质数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func diagonalPrime(nums [][]int) int {
	ans := 0
	for i, m := range nums {
		if m[i] > ans && isPrime(m[i]) {
			ans = m[i]
		}
		if m[len(nums)-1-i] > ans && isPrime(m[len(nums)-1-i]) {
			ans = m[len(nums)-1-i]
		}
	}
	return ans
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[5,6,7],[9,10,11]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[5,17,7],[9,11,10]]\n
// @lcpr case=end

*/

