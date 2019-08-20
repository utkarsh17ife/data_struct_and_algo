package main

import (
	"fmt"
)

type MonthLinkedList struct {
	data string
	next *MonthLinkedList
}

var Head *MonthLinkedList

func main() {
	create()
	read()
	insert()
	read()
}

func create() {

	fmt.Println("Creating linked list")

	firstSixMonths := [6]string{"Jan", "Feb", "Mar", "Apr", "May", "June"}

	var currentNode *MonthLinkedList

	for _, v := range firstSixMonths {

		mll := &MonthLinkedList{
			data: v,
			next: nil,
		}

		if currentNode == nil {
			currentNode = mll
			Head = mll
		} else {
			currentNode.next = mll
			currentNode = mll
		}
	}

}

func read() {

	fmt.Println("Reading linked list")

	var node *MonthLinkedList

	node = Head

	for {
		fmt.Println(node.data)

		if node.next == nil {
			break
		}

		node = node.next
	}

}


func getLastNode() *MonthLinkedList {

	node := Head

	for {

		if node.next == nil {
			return node			
		}

		node = node.next
	}


}

func insert() {

	fmt.Println("inserting into the linked list")

	lastFiveMonths := [6]string{"July", "Aug", "Sept", "Oct", "Nov", "Dec"}
	
	node := getLastNode()

	// treverse to the last node

	for _, v := range lastFiveMonths {

		mll := &MonthLinkedList{
			data: v,
			next: nil,
		}
		
		node.next = mll

		node = node.next
	}

}
