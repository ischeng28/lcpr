/*
 * @lc app=leetcode.cn id=2521 lang=golang
 * @lcpr version=30204
 *
 * [2521] 数组乘积中的不同质因数数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func distinctPrimeFactors(nums []int) int {
	set := map[int]struct{}{}
	for _, x := range nums {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				set[i] = struct{}{}
				for x /= i; x%i == 0; x /= i {
				}
			}
		}
		if x > 1 {
			set[x] = struct{}{}
		}
	}
	return len(set)
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,3,7,10,6]\n
// @lcpr case=end

// @lcpr case=start
// [2,4,8,16]\n
// @lcpr case=end

*/

