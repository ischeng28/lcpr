/*
 * @lc app=leetcode.cn id=1423 lang=golang
 * @lcpr version=30204
 *
 * [1423] 可获得的最大点数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	sum := 0
	res := math.MaxInt32
	total := 0
	for _, v := range cardPoints {
		total += v
	}
	if n == k {
		return total
	}

	for i := 0; i < n; i++ {
		sum += cardPoints[i]
		if i < n-k-1 {
			continue
		}
		res = min(sum, res)
		sum -= cardPoints[i-n+k+1]
	}
	return total - res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2]\n2\n
// @lcpr case=end

// @lcpr case=start
// [9,7,7,9,7,7,9]\n7\n
// @lcpr case=end

// @lcpr case=start
// [1,1000,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,79,80,1,1,1,200,1]\n3\n
// @lcpr case=end

*/

