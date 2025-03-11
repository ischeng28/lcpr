/*
 * @lc app=leetcode.cn id=2105 lang=golang
 * @lcpr version=30204
 *
 * [2105] 给植物浇水 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumRefill(plants []int, capacityA int, capacityB int) int {
	n := len(plants)
	left, right := 0, n-1
	count := 0
	leftRemain, rightRemain := capacityA, capacityB
	for left < right {
		if leftRemain < plants[left] {
			count++
			leftRemain = capacityA
		}
		leftRemain -= plants[left]
		left++
		if rightRemain < plants[right] {
			count++
			rightRemain = capacityB
		}
		rightRemain -= plants[right]
		right--

	}
	if left == right && max(leftRemain, rightRemain) < plants[right] {
		count++
	}
	return count
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,3,3]\n5\n5\n
// @lcpr case=end

// @lcpr case=start
// [2,2,3,3]\n3\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n10\n8\n
// @lcpr case=end

*/

