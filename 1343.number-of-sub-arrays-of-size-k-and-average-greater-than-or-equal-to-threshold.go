/*
 * @lc app=leetcode.cn id=1343 lang=golang
 * @lcpr version=30204
 *
 * [1343] 大小为 K 且平均值大于等于阈值的子数组数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numOfSubarrays(arr []int, k int, threshold int) int {
	res := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i < k-1 {
			continue
		}
		if sum >= threshold*k {
			res++
		}
		sum -= arr[i-k+1]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,2,2,5,5,5,8]\n3\n4\n
// @lcpr case=end

// @lcpr case=start
// [11,13,17,23,29,31,7,5,2,3]\n3\n5\n
// @lcpr case=end

*/

