/*
 * @lc app=leetcode.cn id=1346 lang=golang
 * @lcpr version=30204
 *
 * [1346] 检查整数及其两倍数是否存在
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for i, v := range arr[:len(arr)-1] {
		index := SearchInts(arr, 2*v)
		if index > -1 && arr[index] == 2*v && index != i {
			return true
		}
	}
	return false
}

func SearchInts(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

// @lc code=end

/*
// @lcpr case=start
// [10,2,5,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,1,14,11]\n
// @lcpr case=end

// @lcpr case=start
// [3,1,7,11]\n
// @lcpr case=end

*/

