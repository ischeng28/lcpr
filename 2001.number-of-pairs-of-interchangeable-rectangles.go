/*
 * @lc app=leetcode.cn id=2001 lang=golang
 * @lcpr version=30204
 *
 * [2001] 可互换矩形的组数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func interchangeableRectangles(rectangles [][]int) int64 {
	mp := map[[2]int]int64{}
	var res int64
	for _, v := range rectangles {
		m := gcd(v[0], v[1])
		if count, ok := mp[[2]int{v[0] / m, v[1] / m}]; ok {
			res += count
		}
		mp[[2]int{v[0] / m, v[1] / m}]++
	}
	return res
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [[4,8],[3,6],[10,20],[15,30]]\n
// @lcpr case=end

// @lcpr case=start
// [[4,5],[7,8]]\n
// @lcpr case=end

*/

