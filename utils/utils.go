package utils

import (
	"fmt"
	"os"
	"strconv"
)

type Coordinates struct {
	X int
	Y int
}

type Ants struct {
	AntNum int
	Rooms  []string
	// Position or should we change Rooms to a map of coordinates to room names?
	Position     []Coordinates
	Tunnels      map[string][]string
	StartingRoom string
	EndingRoom   string
}

// The rooms names will not necessarily be numbers, and in order.
// The rooms are identified by a string. it could be "A", "B", "1", "2", etc.  <--- rooms are not necessarily numbers

type DoublyLinkedList struct {
	Head *Node // Pointer to the first node
	Tail *Node // Pointer to the last node
}
type Room struct {
	RoomName string
	AntExist bool
}

// doubly linked list so we can check if there are ants in prev room and the next
type Node struct {
	Room *Room // Data (ant exists?)
	Next *Node // Pointer to the next node
	Prev *Node // Pointer to the previous node
}

func (list *DoublyLinkedList) AddRoom(roomName string) {
	newRoom := &Room{RoomName: roomName, AntExist: false}
	newNode := &Node{Room: newRoom}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
		list.Tail = newNode
	}
}

func Atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("ERROR: invalid data format, Make sure the Coordinates are numbers")
		os.Exit(0)
	}
	return val
}
