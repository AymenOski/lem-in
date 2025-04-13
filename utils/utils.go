package utils

import (
	"fmt"
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

type Ant struct {
	Id          string
	CurrentRoom *Room
}

type Room struct {
	Name      string
	Occupied  bool
	Neighbors []*Room
}

type Graph struct {
	Rooms map[string]*Room
	Paths [][]string
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

func (g *Graph) BFS(start, end string, ant *Ant) []string {
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
			if !visited[neighbor.Name] && !neighbor.Occupied {
				visited[neighbor.Name] = true
				prev[neighbor.Name] = current
				queue = append(queue, neighbor.Name)
			}
		}
	}

	var path []string
	for at := end; at != ""; at = prev[at] {
		if at != end {
			g.Rooms[at].Occupied = true
		}
		path = append([]string{at}, path...)
	}
	return path
}

func (g *Graph) Simulation(ants []*Ant) {
}
