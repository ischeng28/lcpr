/*
 * @lc app=leetcode.cn id=2080 lang=golang
 * @lcpr version=30204
 *
 * [2080] 区间内查询数字的频率
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type RangeFreqQuery map[int][]int

func Constructor(arr []int) RangeFreqQuery {
	mp := map[int][]int{}
	for i, v := range arr {
		mp[v] = append(mp[v], i)
	}
	return mp
}

func (pos RangeFreqQuery) Query(left, right, value int) int {
	a := pos[value]
	return SearchInts(a, right+1) - SearchInts(a, left)

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

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
// @lc code=end

/*
// @lcpr case=start
// ["RangeFreqQuery", "query", "query"][[[12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56]], [1, 2, 4], [0, 11, 33]]\n
// @lcpr case=end

*/

