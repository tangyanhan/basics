/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* head = NULL;
        ListNode* p = NULL;
        ListNode* ia = l1;
        ListNode* ib = l2;
        int addition = 0;
        while(ia || ib || addition) {
            int a=0;
            int b=0;
            if(ia) {
                a = ia->val;
                ia = ia->next;
            }
            if(ib) {
                b = ib->val;
                ib = ib->next;
            }
            
            int val = a + b + addition;
            if( val >= 10 ) {
                val %= 10;
                addition = 1;
            }else{
                addition = 0;
            }
            
            if(p) {
                p->next = new ListNode(val);
                p = p->next;
            }else{
                p = new ListNode(val);
                head = p;
            }
        }
        return head;
    }
};
