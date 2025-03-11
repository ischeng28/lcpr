/*
 * @lc app=leetcode.cn id=1574 lang=golang
 * @lcpr version=30204
 *
 * [1574] 删除最短的子数组使剩余数组有序
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findLengthOfShortestSubarray(arr []int) (ans int) {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 {
		return 0
	}
	ans = right
	left := 0
	for {
		for right == n || arr[left] <= arr[right] {
			ans = min(ans, right-left-1)
			if arr[left] > arr[left+1] {
				return ans
			}
			left++
		}
		right++
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,10,4,2,3,5]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

