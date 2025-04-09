/*
 * @lc app=leetcode.cn id=3115 lang=golang
 * @lcpr version=30204
 *
 * [3115] 质数的最大距离
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
const mx = 101

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

func maximumPrimeDifference(nums []int) int {
	i := 0
	for notPrime[nums[i]] {
		i++
	}
	j := len(nums) - 1
	for notPrime[nums[j]] {
		j--
	}
	return j - i
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,9,5,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,8,2,8]\n
// @lcpr case=end

*/

