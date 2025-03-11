/*
 * @lc app=leetcode.cn id=219 lang=golang
 * @lcpr version=30204
 *
 * [219] 存在重复元素 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	mp := map[int]int{}
	for i, v := range nums {
		if j, ok := mp[v]; ok {
			if i-j <= k {
				return true
			}
		}
		mp[v] = i
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,2,3]\n2\n
// @lcpr case=end

*/

