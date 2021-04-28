package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	linkList := createLinkedList(10)
	linkList.print()

	var index uint = 2
	node := linkList.findByIndex(index)
	fmt.Printf("findIndex %v %v\n", index, node.data)
}

//头指针，不存储数据，单纯指向头结点(有的话)/没有头结点第一个元素，必须有
//头结点，可以作为统计链表数据的节点，可有可无
//也就是有可能链表的第三个位置才是第一个实际的元素(有头结点)

type listNode struct {
	next *listNode
	data interface{}
}

type linkedList struct {
	head   *listNode
	length uint
}

func newListNode(v interface{}) *listNode {
	return &listNode{nil, v}
}

func newLinkedList() *linkedList {
	return &linkedList{
		head:   nil,
		length: 0,
	}
}

//创建一个新的链表
//头指针->第一个元素->第二个元素
func createLinkedList(n int) *linkedList {
	l := newLinkedList()
	r := new(listNode)               //指向最后一个节点
	rand.Seed(time.Now().UnixNano()) //这里不种种子，每次打印出来的数据都是一样的
	if l.length == 0 {               //主要因为链表head这个字段导致
		p := newListNode(rand.Intn(100))
		l.head = p
		fmt.Println(p.data)
		r = p
		l.length++
	}

	for i := 1; i < n; i++ {
		p := newListNode(rand.Intn(100))
		r.next = p
		l.length++
		r = p
	}
	r.next = nil
	return l
}

//在某个节点前插入元素
//前面插入因此必须找到该节点前面的节点
func (l *linkedList) insertBefore(p *listNode, v interface{}) bool {
	if p == nil {
		return false
	}
	cur := l.head
	if cur == p {
		p.next = cur
		l.head = p
		l.length++
		return true
	}
	pre := new(listNode)
	for cur != nil {
		pre = cur
		cur = cur.next
		if cur == p {
			break
		}
	}
	newNode := newListNode(v)
	newNode.next = pre.next
	pre.next = newNode

	return true
}

//某个节点后面插入
func (l *linkedList) insertAfter(p *listNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := newListNode(v)
	newNode.next = p.next
	p.next = newNode
	l.length++
	return true

}

//寻找第n个元素
func (l *linkedList) findByIndex(index uint) *listNode {
	if index > l.length {
		return nil
	}
	var i uint
	cur := l.head
	for cur != nil {
		i++ //注意有头指针的存在
		if i == index {
			return cur
		}
		cur = cur.next
	}
	return nil
}

//打印链表
func (l *linkedList) print() {
	format := ""
	cur := l.head
	for cur != nil {
		format += fmt.Sprintf("val+%v ", cur.data)
		cur = cur.next
	}
	fmt.Println(format)
}
