/*
 * @lc app=leetcode.cn id=744 lang=golang
 * @lcpr version=30204
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := -1, len(letters)
	for left+1 < right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	if right == len(letters) {
		return letters[0]
	}
	return letters[right]
}

// @lc code=end

/*
// @lcpr case=start
// ["c", "f"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// ["c","f","j"]\n"c"\n
// @lcpr case=end

// @lcpr case=start
// ["x","x","y","y"]\n"z"\n
// @lcpr case=end

*/

