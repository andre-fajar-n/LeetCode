/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 != nil { return list2 }
    if list1 != nil && list2 == nil { return list1 }
    if list1 == nil && list2 == nil { return nil }

    output := new(ListNode)
    temp := output
    // i:=0
    // for {
    //     fmt.Println("CEK", list1.Val)
    //     list1 = list1.Next
    //     i++
    //     if list1 == nil {
    //         break
    //     }
    // }
    for list1 != nil || list2 != nil {
        if list1 != nil && list2 != nil {
            if list2.Val <= list1.Val {
                temp.Val = list2.Val
                list2 = list2.Next
            } else {
                temp.Val = list1.Val
                list1 = list1.Next
            }
        } else if list1 == nil && list2 != nil {
            temp.Val = list2.Val
            list2 = list2.Next
        } else if list1 != nil && list2 == nil {
            temp.Val = list1.Val
            list1 = list1.Next
        }
        temp.Next = new(ListNode)
        if list1 == nil && list2 == nil {
            temp.Next = nil
        }
        temp = temp.Next
    }
    // minValue := list1.Val
    // maxValue := list2.Val
    // if list2.Val <= list1.Val {
    //     minValue = list2.Val
    //     maxValue = list1.Val
    // }
    // temp.Val = minValue
    // temp.Next = &ListNode{Val:maxValue, Next: nil}
    // temp = nil

    fmt.Printf(" %+v\n", temp)
    return output
}