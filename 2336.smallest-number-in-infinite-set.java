/*
 * @lc app=leetcode.cn id=2336 lang=java
 * @lcpr version=30204
 *
 * [2336] 无限集中的最小数字
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class SmallestInfiniteSet {

    private int minValue;
    private TreeSet<Integer> set;

    public SmallestInfiniteSet() {
        minValue = 1;
        set = new TreeSet<>();
    }

    public int popSmallest() {
        if (!set.isEmpty()) {
            return set.pollFirst();
        }
        int res = minValue;
        minValue++;
        return res;
    }

    public void addBack(int num) {
        if (num < minValue) {
            set.add(num);
        }
    }
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * SmallestInfiniteSet obj = new SmallestInfiniteSet();
 * int param_1 = obj.popSmallest();
 * obj.addBack(num);
 */
// @lc code=end
