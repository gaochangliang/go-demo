package _1_linkedlist

import "fmt"

type listNode struct {
	next  *listNode
	value interface{}
}

type linkedList struct {
	head *listNode
}

//print link list
func (l *linkedList) Print() {
	cur := l.head
	if cur != nil {
		fmt.Println("value -", cur.value)
		cur = cur.next
	}
}

//reverse link list
func (l *linkedList) reverse() {
	if l.head == nil || l.head.next == nil {
		return
	}

	tmp := l.head
	cur := l.head

	for l.head.next != nil {
		cur = l.head.next
		l.head.next = cur.next //先要保证关系的建立，然后才能移动节点
		cur.next = tmp
		tmp = cur
	}
	l.head = cur
}

//Determining whether a single linked list  has a ring
func (l *linkedList) HasCycle() bool {
	if l.head != nil {
		slow := l.head
		fast := l.head
		//In the case of closed loop, surely the last element cannot be nil
		//using fast.next.next, neither fast nor fast.next can be nil
		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

//
//注意第一次循环后fast是到第三个结点不是第二个结点，slow也是到第二个结点
func (l *linkedList) findMiddleNode() *listNode {
	if l.head == nil {
		return nil
	}

	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

//删除倒数第n个元素
func (l *linkedList) DeleteBottomN(n int) {
	if n < 0 || l.head == nil {
		return
	}

	fast := l.head
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.next
	}

	slow := l.head
	for fast != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next

}


