package utils

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var Mu sync.Mutex

type Coordinates struct {
	X int
	Y int
}

type Ants struct {
	AntNum       int
	Rooms        []string
	Position     []Coordinates
	Tunnels      map[string][]string
	StartingRoom string
	EndingRoom   string
}

type Room struct {
	Name      string
	Visited   bool
	Neighbors []*Room
}

type Graph struct {
	Rooms map[string]*Room
}

func (g *Graph) AddRoom(roomName string) {
	if _, exists := g.Rooms[roomName]; !exists {
		g.Rooms[roomName] = &Room{
			Name:      roomName,
			Neighbors: make([]*Room, 0),
		}
	}
}

func (g *Graph) LinkRooms(tunnels map[string][]string) {
	for roomName, neighbors := range tunnels {
		for _, neighborName := range neighbors {
			if g.Rooms[roomName] != nil && g.Rooms[neighborName] != nil {
				g.Rooms[roomName].Neighbors = append(g.Rooms[roomName].Neighbors, g.Rooms[neighborName])
			}
		}
	}
}

func (g *Graph) PrintGraph() {
	for _, room := range g.Rooms {
		fmt.Printf("Room: %s, Neighbors: %v\n", room.Name, room.Neighbors)
	}
}

func (g *Graph) BFS(start, end string, AntNum int) {
	c := 1
	temp := strconv.Itoa(c)
	str := fmt.Sprintf("L" + temp + "-")

	queue := []string{start}
	g.Rooms[start].Visited = true
	prev := make(map[string]string)

	for len(queue) > 0 {

		current := queue[0]
		// deleat from queue
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range g.Rooms[current].Neighbors {
			if !neighbor.Visited {
				neighbor.Visited = true
				prev[neighbor.Name] = current
				queue = append(queue, neighbor.Name)
			}
		}
	}

	var path []string
	for at := end; at != start; at = prev[at] {
		path = append([]string{at}, path...)
	}
	path = append([]string{start}, path...)
	if len(path) == 0 || path[0] != start {
		fmt.Println("No path found")
		os.Exit(0)
	}
	// matrix := [][]string{}
	for _, room := range path[1:] {
		str += room
		fmt.Println(str)
		str = str[:len(str)-len(room)]
	}
	path = nil
	AntNum--
	c++
}
