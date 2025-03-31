/*
 * @lc app=leetcode.cn id=2187 lang=golang
 * @lcpr version=30204
 *
 * [2187] 完成旅途的最少时间
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumTime(time []int, totalTrips int) int64 {
	check := func(m int) bool {
		sum := 0
		for _, v := range time {
			sum += m / v
		}
		return sum >= totalTrips
	}

	left, right := 0, slices.Max(time)*totalTrips+1
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return int64(right)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/

