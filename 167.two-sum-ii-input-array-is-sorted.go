/*
 * @lc app=leetcode.cn id=167 lang=golang
 * @lcpr version=30204
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n-1
	for left < right {
		tmp := numbers[left] + numbers[right]
		if tmp == target {
			return []int{left + 1, right + 1}
		}
		if tmp < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [-1,0]\n-1\n
// @lcpr case=end

*/

