/*
 * @lc app=leetcode.cn id=1052 lang=golang
 * @lcpr version=30204
 *
 * [1052] 爱生气的书店老板
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	sum1 := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			sum1 += customers[i]
		}
	}
	sum2 := 0
	maxSum2 := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 1 {
			sum2 += customers[i]
		}
		if i < minutes-1 {
			continue
		}
		maxSum2 = max(maxSum2, sum2)
		if grumpy[i-minutes+1] == 1 {
			sum2 -= customers[i-minutes+1]
		}
	}
	return maxSum2 + sum1
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,1,2,1,1,7,5]\n[0,1,0,1,0,1,0,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n[0]\n1\n
// @lcpr case=end

*/

