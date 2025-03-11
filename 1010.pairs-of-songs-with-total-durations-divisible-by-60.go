/*
 * @lc app=leetcode.cn id=1010 lang=golang
 * @lcpr version=30204
 *
 * [1010] 总持续时间可被 60 整除的歌曲
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numPairsDivisibleBy60(time []int) int {
	res := 0
	mp := map[int]int{}
	for _, v := range time {
		if j, ok := mp[(60-v%60)%60]; ok {
			res = res + j
		}
		mp[v%60]++
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [30,20,150,100,40]\n
// @lcpr case=end

// @lcpr case=start
// [60,60,60]\n
// @lcpr case=end

*/

