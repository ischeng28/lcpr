/*
 * @lc app=leetcode.cn id=1870 lang=golang
 * @lcpr version=30204
 *
 * [1870] 准时到达的列车最小时速
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minSpeedOnTime(dist []int, hour float64) int {
	if hour <= float64(len(dist)-1) {
		return -1
	}
	check := func(m int) bool {
		sum := 0
		for _, v := range dist[:len(dist)-1] {
			sum += (v + m - 1) / m
		}
		return float64(sum)+float64(dist[len(dist)-1])/float64(m) <= hour
	}
	left, right := 0, slices.Max(dist)*100+1
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
// [1,3,2]\n6\n
// @lcpr case=end

// @lcpr case=start
// [1,3,2]\n2.7\n
// @lcpr case=end

// @lcpr case=start
// [1,3,2]\n1.9\n
// @lcpr case=end

*/

