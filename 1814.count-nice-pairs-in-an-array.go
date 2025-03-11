/*
 * @lc app=leetcode.cn id=1814 lang=golang
 * @lcpr version=30204
 *
 * [1814] 统计一个数组中好对子的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countNicePairs(nums []int) (res int) {
	mp := map[int]int{}
	for _, v := range nums {
		if e, ok := mp[rev(v)-v]; ok {
			res += e
		}
		mp[rev(v)-v]++
	}
	return res % (1e9 + 7)
}

func rev(num int) (res int) {
	for num > 0 {
		res = res*10 + num%10
		num = num / 10
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [42,11,1,97]\n
// @lcpr case=end

// @lcpr case=start
// [13,10,35,24,76]\n
// @lcpr case=end

*/

