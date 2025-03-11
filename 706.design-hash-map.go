/*
 * @lc app=leetcode.cn id=706 lang=golang
 * @lcpr version=30204
 *
 * [706] 设计哈希映射
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

const base = 769

type entry struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{make([]list.List,base)}
}

func ()

func (this *MyHashMap) Put(key int, value int) {

}

func (this *MyHashMap) Get(key int) int {

}

func (this *MyHashMap) Remove(key int) {

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

/*
// @lcpr case=start
// ["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"][[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]\n
// @lcpr case=end

*/
