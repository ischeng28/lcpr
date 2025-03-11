/*
 * @lc app=leetcode.cn id=278 lang=golang
 * @lcpr version=30204
 *
 * [278] 第一个错误的版本
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left, right := 0, n+1
	for left+1 < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
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
// 5\n4\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/

