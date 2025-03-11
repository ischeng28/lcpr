/*
 * @lc app=leetcode.cn id=374 lang=golang
 * @lcpr version=30204
 *
 * [374] 猜数字大小
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 0, n+1
	for left+1 < right {
		mid := left + (right-left)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) > 0 {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// 10\n6\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

// @lcpr case=start
// 2\n1\n
// @lcpr case=end

*/

