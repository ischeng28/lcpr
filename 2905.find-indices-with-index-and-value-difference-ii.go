/*
 * @lc app=leetcode.cn id=2905 lang=golang
 * @lcpr version=30204
 *
 * [2905] 找出满足差值条件的下标 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	maxPre, minPre := 0, 0
	for j := indexDifference; j < len(nums); j++ {
		i := j - indexDifference
		if nums[i] > nums[maxPre] {
			maxPre = i
		} else if nums[i] < nums[minPre] {
			minPre = i
		}
		if nums[maxPre]-nums[j] >= valueDifference {
			return []int{maxPre, j}
		}
		if nums[j]-nums[minPre] >= valueDifference {
			return []int{minPre, j}
		}
	}
	return []int{-1, -1}
}

// @lc code=end

/*
// @lcpr case=start
// [5,1,4,1]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n0\n0\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n2\n4\n
// @lcpr case=end

*/

func findIndices(nums []int, indexDifference, valueDifference int) []int {
	maxIdx, minIdx := 0, 0
	for j := indexDifference; j < len(nums); j++ {
		i := j - indexDifference
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		} else if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[maxIdx]-nums[j] >= valueDifference {
			return []int{maxIdx, j}
		}
		if nums[j]-nums[minIdx] >= valueDifference {
			return []int{minIdx, j}
		}
	}
	return []int{-1, -1}
}


