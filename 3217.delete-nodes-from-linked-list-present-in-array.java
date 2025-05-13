/*
 * @lc app=leetcode.cn id=3217 lang=java
 * @lcpr version=30204
 *
 * [3217] 从链表中移除在数组中存在的节点
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.Set;

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
    public ListNode modifiedList(int[] nums, ListNode head) {
        Set<Integer> set = new HashSet<>();
        for (int num : nums) {
            set.add(num);
        }
        ListNode dummy = new ListNode(0, head);
        ListNode current = dummy;
        while (current.next != null) {
            if (set.contains(current.next.val)) {
                current.next = current.next.next;
            } else {
                current = current.next;
            }
        }
        return dummy.next;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n[1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n[1,2,1,2,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [5]\n[1,2,3,4]\n
// @lcpr case=end

 */
