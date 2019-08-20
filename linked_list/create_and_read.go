package main


import (
	"fmt"
)

type MonthLinkedList struct {
	data string
	next *MonthLinkedList
}

var Head *MonthLinkedList

func main (){
	create()
	read()
}

func create() {

	fmt.Println("Creating linked list")

	months := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"}

	var currentNode *MonthLinkedList

	for _, v := range months {

		mll := &MonthLinkedList{
			data: v,
			next: nil,
		}

		if currentNode == nil {
			currentNode = mll
			Head = mll
		}else {
			currentNode.next = mll
			currentNode = mll
		}
	}

}

func read(){

	fmt.Println("Reading linked list")

	var node *MonthLinkedList

	node = Head

	for {
		fmt.Println(node.data)		

		if node.next == nil {
			break;
		}

		node = node.next
	}

}
