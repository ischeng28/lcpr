/*
 * @lc app=leetcode.cn id=933 lang=java
 * @lcpr version=30204
 *
 * [933] 最近的请求次数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.Queue;

class RecentCounter {
    private Queue<Integer> queue;

    public RecentCounter() {
        queue = new ArrayDeque<Integer>();
    }

    public int ping(int t) {
        queue.offer(t);
        while (queue.peek() < t - 3000) {
            queue.poll();
        }
        return queue.size();
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * RecentCounter obj = new RecentCounter();
 * int param_1 = obj.ping(t);
 */
// @lc code=end

/*
 * // @lcpr case=start
 * // ["RecentCounter", "ping", "ping", "ping", "ping"][[], [1], [100], [3001],
 * [3002]]\n
 * // @lcpr case=end
 * 
 */
