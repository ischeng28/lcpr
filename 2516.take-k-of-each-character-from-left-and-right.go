/*
 * @lc app=leetcode.cn id=2516 lang=golang
 * @lcpr version=30204
 *
 * [2516] 每种字符至少取 K 个
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func takeCharacters(s string, k int) int {
	arr := make([]int, 3)
	// 计数的同时全部拿走，之后再通过窗口滑动，挨个放回来，这样能够实现数组的值是取走的个数的效果
	for _, v := range s {
		arr[v-'a']++
	}
	if arr[0] < k || arr[1] < k || arr[2] < k {
		return -1
	}
	left := 0
	length := 0
	for right, v := range s {
		index := v - 'a'
		arr[index]--
		for arr[index] < k {
			arr[s[left]-'a']++
			left++
		}
		length = max(length, right-left+1)
	}
	return len(s) - length
}

// @lc code=end

/*
// @lcpr case=start
// "aabaaaacaabc"\n2\n
// @lcpr case=end

// @lcpr case=start
// "a"\n1\n
// @lcpr case=end

*/

