/*
 * @lc app=leetcode.cn id=1669 lang=java
 * @lcpr version=30204
 *
 * [1669] 合并两个链表
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
    public ListNode mergeInBetween(ListNode list1, int a, int b, ListNode list2) {
        ListNode p = list1;
        for (int i = 1; i < a; i++) {
            p = p.next;
        }
        ListNode remove = p.next;
        p.next = list2;
        while (p.next != null) {
            p = p.next;
        }
        for (int i = a; i <= b; i++) {
            remove = remove.next;
        }
        p.next = remove;
        return list1;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [10,1,13,6,9,5]\n3\n4\n[1000000,1000001,1000002]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,3,4,5,6]\n2\n5\n[1000000,1000001,1000002,1000003,1000004]\n
// @lcpr case=end

 */
