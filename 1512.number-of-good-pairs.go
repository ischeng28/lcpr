/*
 * @lc app=leetcode.cn id=1512 lang=golang
 * @lcpr version=30204
 *
 * [1512] 好数对的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numIdenticalPairs(nums []int) int {
	res := 0
	mp := map[int]int{}
	for _, v := range nums {
		if j, ok := mp[v]; ok {
			res = res + j
		}
		mp[v]++
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

