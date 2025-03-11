/*
 * @lc app=leetcode.cn id=27 lang=golang
 * @lcpr version=30204
 *
 * [27] 移除元素
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func removeElement(nums []int, val int) int {
	left := 0
	count := 0
	for right, v := range nums {
		if v == val {
			continue
		}
		nums[left] = nums[right]
		left++
		count++
	}
	return count
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2,3,0,4,2]\n2\n
// @lcpr case=end

*/

