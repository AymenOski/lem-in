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
	Id          string
	CurrentRoom *Room
	Path        []string
	HasMoved    bool
	Step        int
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

func (g *Graph) BFS(start, end string) []string {
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	prev := make(map[string]string)
	for len(queue) >= 0 {
		fmt.Println("ss")

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
	if !visited[end] {
		fmt.Println("kk")
		return nil
	}
	var path []string
	for at := end; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
		if at != start && at != end {
			g.Rooms[at].Occupied = true
		}
	}
	return path
}

func (g *Graph) Simulation(ants []*Ant, Start string, End string) {
	pathLens := make([]int, len(g.Paths))

	for _, room := range g.Rooms {
		room.Occupied = false
	}

	for i := range pathLens {
		pathLens[i] = len(g.Paths[i])
	}
	// Assign paths to ants
	for _, ant := range ants {
		
		// we are implementing a !(greedy algorithm) to assign what is best for the greater good of the colony
		bestIdx := 0
		// we pick the shortest path based on lenght
		for i := 1; i < len(pathLens); i++ {
			if pathLens[bestIdx] > pathLens[i] {
				bestIdx = i
			}
		}
		
		ant.Path = g.Paths[bestIdx]
		
		pathLens[bestIdx]++
	}
	allReachedEnd := false

	for !allReachedEnd {
		for _, ant := range ants {
			if ant.CurrentRoom != g.Rooms[End] && ant.Step < len(ant.Path)-1 {
				nextRoomName := ant.Path[ant.Step+1]
				nextRoom := g.Rooms[nextRoomName]

				if !nextRoom.Occupied || nextRoom.Name == End {
					ant.CurrentRoom = nextRoom
					ant.CurrentRoom.Occupied = true

					ant.Step++
					ant.HasMoved = true
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

		allReachedEnd = true
		for _, ant := range ants {
			if ant.CurrentRoom != g.Rooms[End] {
				allReachedEnd = false
			}
		}

	}

	// for _, ant := range ants {
	// 	fmt.Println(ant.Path)
	// }
}
