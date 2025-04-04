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
	AntNum       int
	Rooms        []string
	Position     []Coordinates
	Tunnels      map[string][]string
	StartingRoom string
	EndingRoom   string
}

// The rooms names will not necessarily be numbers, and in order.
// The rooms are identified by a string. it could be "A", "B", "1", "2", etc.  <--- rooms are not necessarily numbers

type Room struct {
	Name      string
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
			g.Rooms[roomName].Neighbors = append(g.Rooms[roomName].Neighbors, g.Rooms[neighborName])
		}
	}
}

func (g *Graph) PrintGraph() {
	for _, room := range g.Rooms {
		fmt.Printf("Room: %s, Neighbors: %v\n", room.Name, room.Neighbors)
	}
}

func (g *Graph) BFS(start, end string) []string {
	queue := []string{start}

	visited := make(map[string]bool)
	visited[start] = true

	prev := make(map[string]string)

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range g.Rooms[current].Neighbors {
			if !visited[neighbor.Name] {
				visited[neighbor.Name] = true
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
		return nil
	}

	return path
}

func Atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("ERROR: invalid data format, Make sure the Coordinates are numbers")
		os.Exit(0)
	}
	return val
}
