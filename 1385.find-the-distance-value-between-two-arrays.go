/*
 * @lc app=leetcode.cn id=1385 lang=golang
 * @lcpr version=30204
 *
 * [1385] 两个数组间的距离值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	res := 0
	for _, v := range arr1 {
		low := v - d
		high := v + d
		if !SearchBetween(arr2, low, high) {
			res++
		}
	}
	return res
}

func SearchBetween(nums []int, low int, high int) bool {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] >= low && nums[mid] <= high {
			return true
		}
		if nums[mid] < low {
			left = mid
		} else {
			right = mid
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [4,5,8]\n[10,9,1,8]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,4,2,3]\n[-4,-3,6,10,20,30]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,1,100,3]\n[-5,-2,10,-3,7]\n6\n
// @lcpr case=end

*/

