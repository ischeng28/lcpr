/*
 * @lc app=leetcode.cn id=658 lang=golang
 * @lcpr version=30204
 *
 * [658] 找到 K 个最接近的元素
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1
	for right-left+1 > k {
		if abs(arr[right]-x) >= abs(x-arr[left]) {
			right--
		} else {
			left++
		}
	}
	return arr[left : right+1]
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n4\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,3,4,5]\n4\n-1\n
// @lcpr case=end

*/

