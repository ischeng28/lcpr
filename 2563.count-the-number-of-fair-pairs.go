/*
 * @lc app=leetcode.cn id=2563 lang=golang
 * @lcpr version=30204
 *
 * [2563] 统计公平数对的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countFairPairs(nums []int, lower int, upper int) int64 {
	ans := 0
	if len(nums) == 1 {
		if nums[0] >= lower && nums[0] <= upper {
			return 1
		}
		return 0
	}
	sort.Ints(nums)
	for i, v := range nums {
		arr := nums[i+1:]
		r := SearchInts(arr, upper-v+1)
		l := SearchInts(arr, lower-v)
		ans += r - l
	}
	return int64(ans)

}

func SearchInts(arr []int, target int) int {
	left, right := -1, len(arr)
	for left+1 < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,7,4,4,5]\n3\n6\n
// @lcpr case=end

// @lcpr case=start
// [1,7,9,2,5]\n11\n11\n
// @lcpr case=end

*/

