/*
 * @lc app=leetcode.cn id=2748 lang=golang
 * @lcpr version=30204
 *
 * [2748] 美丽下标对的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countBeautifulPairs(nums []int) int {
	res := 0
	mp := map[int]int{}
	for _, v := range nums {
		for j := 1; j < 10; j++ {
			if gcd(v%10, j) == 1 {
				res += mp[j]
			}
		}
		for v >= 10 {
			v = v / 10
		}
		mp[v]++
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [2,5,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [11,21,12]\n
// @lcpr case=end

*/

