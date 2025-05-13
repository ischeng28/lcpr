/*
 * @lc app=leetcode.cn id=950 lang=java
 * @lcpr version=30204
 *
 * [950] 按递增顺序显示卡牌
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
import java.util.Deque;

class Solution {
    public int[] deckRevealedIncreasing(int[] deck) {
        Arrays.sort(deck);
        Deque<Integer> deque = new ArrayDeque<Integer>();
        for (int i = deck.length - 1; i >= 0; i--) {
            if (!deque.isEmpty()) {
                deque.addFirst(deque.pollLast());
            }
            deque.addFirst(deck[i]);
        }
        int[] res = new int[deck.length];
        for (int k = 0; k < deck.length; k++) {
            System.out.println(deque.peekFirst());
            res[k] = deque.pollFirst();
        }
        return res;
    }
}
// @lc code=end



/*
// @lcpr case=start
// [17,13,11,2,3,5,7]\n
// @lcpr case=end

 */

