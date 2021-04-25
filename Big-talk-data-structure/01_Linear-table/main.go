package main

import (
	"errors"
	"fmt"
)

//模拟链表如何写
func main() {
	var l, l2, l3 linkList
	l.data = 1
	l2.data = 2
	l3.data = 3
	l.next = &l2
	l2.next = &l3
	number, _ := getElem(&l, 2)
	fmt.Println("number", number)
}

type linkList struct {
	data int
	next *linkList
}

//Get the i-th data of the link list

func getElem(list *linkList, i int) (int, error) {
	if list == nil {
		return 0, errors.New("the link list is empty")
	}
	p := list
	var index = 1
	for p != nil && index < i {
		p = list.next
		index++
	}

	if index < i {
		return 0, errors.New("the nth data does not exist") //链表长度不够
	}
	return p.data, nil
}
