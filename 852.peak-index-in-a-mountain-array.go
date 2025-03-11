/*
 * @lc app=leetcode.cn id=852 lang=golang
 * @lcpr version=30204
 *
 * [852] 山脉数组的峰顶索引
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,2,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,10,5,2]\n
// @lcpr case=end

*/

