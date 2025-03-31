/*
 * @lc app=leetcode.cn id=2342 lang=golang
 * @lcpr version=30204
 *
 * [2342] 数位和相等数对的最大和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumSum(nums []int) int {
	mp := map[int]int{}
	res := -1
	for _, v := range nums {
		if m, ok := mp[getSum(v)]; ok {
			res = max(res, v+m)
		}
		mp[getSum(v)] = max(mp[getSum(v)], v)
	}
	return res
}

func getSum(a int) int {
	sum := 0
	for a > 0 {
		sum += a % 10
		a = a / 10
	}
	return sum
}

// @lc code=end

/*
// @lcpr case=start
// [18,43,36,13,7]\n
// @lcpr case=end

// @lcpr case=start
// [10,12,19,14]\n
// @lcpr case=end

*/

