/*
 * @lc app=leetcode.cn id=2216 lang=golang
 * @lcpr version=30204
 *
 * [2216] 美化数组的最少删除数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minDeletion(nums []int) int {
	stack := []int{}
	for _, v := range nums {
		if len(stack)%2 != 0 {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				continue
			}
		}
		stack = append(stack, v)
	}
	return len(nums) - len(stack)
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2,3,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,2,3,3]\n
// @lcpr case=end

*/

