/*
 * @lc app=leetcode.cn id=875 lang=golang
 * @lcpr version=30204
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	check := func(m int) bool {
		count := 0
		for _, v := range piles {
			count += (v + m - 1) / m
		}
		return count <= h
	}

	left, right := 0, slices.Max(piles)
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
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
// [3,6,7,11]\n8\n
// @lcpr case=end

// @lcpr case=start
// [30,11,23,4,20]\n5\n
// @lcpr case=end

// @lcpr case=start
// [30,11,23,4,20]\n6\n
// @lcpr case=end

*/

