/*
 * @lc app=leetcode.cn id=275 lang=golang
 * @lcpr version=30204
 *
 * [275] H 指数 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func hIndex(citations []int) int {
	left, right := -1, len(citations)
	for left+1 < right {
		mid := left + (right-left)/2
		if citations[mid] >= len(citations)-mid {
			right = mid
		} else {
			left = mid
		}
	}
	return len(citations) - right
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,3,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,100]\n
// @lcpr case=end

*/

