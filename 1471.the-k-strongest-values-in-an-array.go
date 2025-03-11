/*
 * @lc app=leetcode.cn id=1471 lang=golang
 * @lcpr version=30204
 *
 * [1471] 数组中的 k 个最强值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	left, right := 0, len(arr)-1
	res := []int{}
	v := arr[(len(arr)-1)/2]
	for left <= right && len(res) < k {
		if arr[right]-v >= v-arr[left] {
			res = append(res, arr[right])
			right--
		} else {
			res = append(res, arr[left])
			left++
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,1,3,5,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [6,7,11,7,6,8]\n5\n
// @lcpr case=end

// @lcpr case=start
// [6,-3,7,2,11]\n3\n
// @lcpr case=end

// @lcpr case=start
// [-7,22,17,3]\n2\n
// @lcpr case=end

*/

