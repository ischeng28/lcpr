/*
 * @lc app=leetcode.cn id=1472 lang=golang
 * @lcpr version=30204
 *
 * [1472] 设计浏览器历史记录
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type BrowserHistory struct {
	his []string
	idx int
}

func Constructor(homepage string) BrowserHistory {
	his := []string{}
	his = append(his, homepage)
	return BrowserHistory{his: his, idx: 0}
}

func (this *BrowserHistory) Visit(url string) {
	this.idx++
	this.his = this.his[:this.idx]
	this.his = append(this.his, url)
}

func (this *BrowserHistory) Back(steps int) string {
	this.idx = max(this.idx-steps, 0)
	return this.his[this.idx]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.idx = min(this.idx+steps, len(this.his)-1)
	return this.his[this.idx]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
// @lc code=end

/*
// @lcpr case=start
// ["BrowserHistory","visit","visit","visit","back","back","forward","visit","forward","back","back"][["leetcode.com"],["google.com"],["facebook.com"],["youtube.com"],[1],[1],[1],["linkedin.com"],[2],[2],[7]]\n
// @lcpr case=end

*/

