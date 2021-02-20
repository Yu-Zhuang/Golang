package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

func main() {
	fmt.Println("Data Type Practice and note")
	// create head node
	head := node{value: 0, next: nil}
	// create linked list
	for i := 1; i < 10; i++ {
		head.addNode(i)
	}
	// show linked list
	head.showLink()
}

func (head *node) addNode(val int) {
	tmp := head
	// create
	newNode := node{value: val, next: nil}
	for {
		if tmp.next == nil {
			tmp.next = &newNode
			return
		}
		tmp = tmp.next
	}
}

func (head node) showLink() {
	for {
		if head.next == nil {
			fmt.Print("[", head.value, "]-|END\n")
			return
		}
		fmt.Print("[", head.value, "]-")
		head = *head.next
	}
}
