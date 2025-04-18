package utils

import (
	"fmt"
)

type Coordinates struct {
	X int
	Y int
}

type Colony struct {
	AntNum       int
	Rooms        []string
	Position     []Coordinates
	Tunnels      map[string][]string
	StartingRoom string
	EndingRoom   string
}

type Ant struct {
	Id           string
	CurrentRoom  *Room
	ThePath      []string
	HasMoved     bool
	HasTakenPath bool
	Step         int
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
		fmt.Printf("Room: %s, Neighbors: %v\n", room.Name , room.Neighbors)
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

func (g *Graph) Simulation(ants []*Ant, Start string, End string) {
	for _, room := range g.Rooms {
		room.Occupied = false
	}

	// Assign paths to ants
	i := 0
	for _, ant := range ants {
		if !ant.HasTakenPath {
			for {
				if i >= len(g.Paths) {
					i = 0 // If we are out of paths , start over and give the shortest path to the ant
				}
				// Skip ants that dont have a path
				if len(g.Paths[i]) == 1 {
					i++
					continue
				}

				// Assign path to the ant
				ant.ThePath = g.Paths[i]
				i++

				ant.HasTakenPath = true
				ant.Step = 0
				break
			}
		}
	}

	allReachedEnd := false

	for !allReachedEnd {

		for _, ant := range ants {

			if ant.CurrentRoom != g.Rooms[End] && ant.Step < len(ant.ThePath) {
				nextRoomName := ant.ThePath[ant.Step+1]
				nextRoom := g.Rooms[nextRoomName]

				if !nextRoom.Occupied || nextRoom.Name == End {
					ant.CurrentRoom = nextRoom
					ant.CurrentRoom.Occupied = true

					ant.Step++
					ant.HasMoved = true
				}
			}

			allReachedEnd = true
			for _, ant := range ants {
				if ant.CurrentRoom != g.Rooms[End] {
					allReachedEnd = false
				}
			}

		}

		for _, ant := range ants {
			if ant.HasMoved {
				fmt.Printf("%s-%s ", ant.Id, ant.CurrentRoom.Name)
				ant.HasMoved = false
				ant.CurrentRoom.Occupied = false
			}
		}
		fmt.Println()

	}

	// for _, ant := range ants {
	// 	fmt.Println(ant.ThePath)
	// }
}
