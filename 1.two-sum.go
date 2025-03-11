/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30204
 *
 * [1] 两数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func twoSum(nums []int, target int) []int {
	mp := map[int]int{}
	for j, v := range nums {
		if i, ok := mp[target-v]; ok {
			return []int{i, j}
		}
		mp[v] = j
	}
	return []int{}
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/

