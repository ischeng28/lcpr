/*
 * @lc app=leetcode.cn id=1909 lang=golang
 * @lcpr version=30204
 *
 * [1909] 删除一个元素使数组严格递增
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func canBeIncreasing(nums []int) bool {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[left+1] {
			left++
		} else if nums[right] > nums[right-1] {
			right--
		} else {
			break
		}
	}
	if left == right {
		return true
	}
	if left != right-1 {
		return false
	}
	if left == 0 || right == len(nums)-1 || nums[left] < nums[right+1] || nums[right] > nums[left-1] {
		return true
	} else {
		return false
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,10,5,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

