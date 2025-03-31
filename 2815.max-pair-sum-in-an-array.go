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
	res := -1
	mp := map[int]int{}
	for _, v := range nums {
		m := getMaxNum(v)
		if n, ok := mp[m]; ok {
			res = max(res, n+v)
		}
		mp[m] = max(mp[m], v)
	}
	return res
}

func getMaxNum(m int) int {
	res := 0
	for m > 0 {
		res = max(res, m%10)
		m = m / 10
	}
	return res
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

