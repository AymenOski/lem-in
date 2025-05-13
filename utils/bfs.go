package utils

// Find one shortest path using BFS.
func (g *Graph) BFS(start, end string) []string {
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	prev := make(map[string]string)
	skipStartEndPath := false

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}

		for _, neighbor := range g.Rooms[current].Neighbors {
			if !skipStartEndPath {
				if hasLengthString(g.Paths) && neighbor.Name == end {
					skipStartEndPath = true
					continue
				}
			}
			if !visited[neighbor.Name] && (!neighbor.Occupied || neighbor.Name == end) {
				visited[neighbor.Name] = true
				prev[neighbor.Name] = current
				queue = append(queue, neighbor.Name)
			}
		}
	}

	if !visited[end] {
		return nil
	}

	var path []string

	for at := end; at != ""; at = prev[at] {
		// this condition is important to find all possible shortests paths
		if prev[at] == start {
			g.Rooms[at].Occupied = true
		}
		path = append([]string{at}, path...)
	}

	return path
}
