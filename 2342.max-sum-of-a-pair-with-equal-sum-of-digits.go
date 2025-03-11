/*
 * @lc app=leetcode.cn id=2342 lang=golang
 * @lcpr version=30204
 *
 * [2342] 数位和相等数对的最大和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumSum(nums []int) int {
	mp := map[int]int{}
	res := -1
	for _, v := range nums {
		numSum := getNumSum(v)
		if j, ok := mp[numSum]; ok {
			res = max(res, j+v)
			mp[numSum] = max(j, v)
		} else {
			mp[numSum] = v
		}
	}
	return res
}

func getNumSum(num int) (res int) {
	for num > 0 {
		res += num % 10
		num = num / 10
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [18,43,36,13,7]\n
// @lcpr case=end

// @lcpr case=start
// [10,12,19,14]\n
// @lcpr case=end

*/

