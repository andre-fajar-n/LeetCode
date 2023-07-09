# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        temp = dummy
        
        num1 = l1
        num2 = l2
        carry = 0
        while num1 is not None or num2 is not None or carry > 0:
            if num1 is not None:
                carry += num1.val
                num1 = num1.next
            
            if num2 is not None:
                carry += num2.val
                num2 = num2.next

            newNode = ListNode(carry%10)
            temp.next = newNode
            temp = temp.next
            carry //= 10
        
        output = dummy.next
        dummy.next = None
        return output