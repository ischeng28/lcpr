/*
 * @lc app=leetcode.cn id=2181 lang=java
 * @lcpr version=30204
 *
 * [2181] 合并零之间的节点
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode mergeNodes(ListNode head) {
        int sum = 0;
        ListNode res = new ListNode();
        ListNode p = res;
        while (head != null) {
            if (head.val == 0) {
                if (sum != 0) {
                    p.next = new ListNode(sum);
                    p = p.next;
                    sum = 0;
                }
            } else {
                sum += head.val;
            }
            head = head.next;
        }
        return res.next;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [0,3,1,0,4,5,2,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,0,2,2,0]\n
// @lcpr case=end

 */
