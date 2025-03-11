/*
 * @lc app=leetcode.cn id=2815 lang=golang
 * @lcpr version=30204
 *
 * [2815] 数组中的最大数对和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSum(nums []int) int {
	mp := map[int]int{}
	maxSum := -1
	for _, v := range nums {
		if i, ok := mp[maxNum(v)]; ok {
			maxSum = max(maxSum, i+v)
		}
		mp[maxNum(v)] = max(mp[maxNum(v)], v)
	}
	return maxSum
}

func maxNum(num int) int {
	maxNum := 0
	for num > 0 {
		maxNum = max(maxNum, num%10)
		num = num / 10
	}
	return maxNum
}

// @lc code=end

/*
// @lcpr case=start
// [51,71,17,24,42]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

*/

