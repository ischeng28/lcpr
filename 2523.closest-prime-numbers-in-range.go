/*
 * @lc app=leetcode.cn id=2523 lang=golang
 * @lcpr version=30204
 *
 * [2523] 范围内最接近的两个质数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
const mx = 1e6 + 1

var notPrime = [mx]bool{true, true}

func init() {
	for i := 2; i*i < mx; i++ {
		if !notPrime[i] {
			for j := i * i; j < mx; j += i {
				notPrime[j] = true
			}
		}
	}
}

func closestPrimes(left int, right int) []int {
	// 收集范围内的所有质数
	primes := []int{}
	for i := left; i <= right; i++ {
		if !notPrime[i] {
			primes = append(primes, i)
		}
	}

	// 如果质数数量少于2个，返回[-1, -1]
	if len(primes) < 2 {
		return []int{-1, -1}
	}

	// 找到最接近的一对质数
	minDiff := int(1e9)
	result := []int{-1, -1}

	for i := 1; i < len(primes); i++ {
		diff := primes[i] - primes[i-1]
		if diff < minDiff {
			minDiff = diff
			result[0], result[1] = primes[i-1], primes[i]
		}
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// 10\n19\n
// @lcpr case=end

// @lcpr case=start
// 4\n6\n
// @lcpr case=end

*/

