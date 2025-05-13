/*
 * @lc app=leetcode.cn id=817 lang=java
 * @lcpr version=30204
 *
 * [817] 链表组件
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
    public int numComponents(ListNode head, int[] nums) {
        Set<Integer> set = new HashSet<>();
        for (int num : nums) {
            set.add(num);
        }
        int res = 0;
        ListNode current = head;
        while (current != null) {
            if (set.contains(current.val)) {
                while (current != null && set.contains(current.val)) {
                    current = current.next;
                }
                res++;
            } else {
                current = current.next;
            }
        }

        return res;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [0,1,2,3]\n[0,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,3,4]\n[0,3,1,4]\n
// @lcpr case=end

 */
