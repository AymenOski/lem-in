package parse

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	constant "lem-in/const"
	"lem-in/utils"
)

// implement the error interface{}
type ErrorMessage struct {
	Msg string
}

func (e *ErrorMessage) Error() string {
	return e.Msg
}

func FileExist(filename string) error {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return &ErrorMessage{Msg: constant.ErrPrefix + "The File " + filename + " Doesn't Exist in The Specifeid Path"}
	}
	if info.IsDir() {
		return &ErrorMessage{Msg: constant.ErrPrefix + "You Have Entered a Directory Path Istead Of a File Path"}
	}
	return nil
}

func FileToGraph(filename string) (*utils.Graph, error) {
	var err error
	err = FileExist(filename)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrFileIssue}
	}
	defer file.Close()
	graph := utils.GraphConstructor()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		err := ParseLine(graph, scanner.Text())
		if err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
		return nil, err
	}

	if graph.Col == nil || graph.Rooms == nil {
		return nil, &ErrorMessage{Msg: constant.ErrPrefix + "The graph is empty"}
	}

	return graph, nil
}

func ParseLine(graph *utils.Graph, line string) error {
	if IsComment(line) || line == "" {
		return nil
	}
	switch graph.Data.Phase {
	case constant.AntsField:
		return ParseAntNumber(graph, line)
	case constant.RoomsField:
		return ParseRooms(graph, line)
	case constant.LinksField:
		return ParseLinks(graph, line)
	default:
		return &ErrorMessage{Msg: constant.ErrPrefix + "something went wrong while parsing"}
	}
}

func ParseAntNumber(graph *utils.Graph, line string) error {
	graph.Col = &utils.Colony{}
	n, err := strconv.Atoi(line)
	if err != nil {
		return &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrAnts}
	}
	if n <= 0 || n > (1<<31-1) {
		return &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrAnts}
	}
	graph.Col.AntNum = n
	graph.Data.Phase = constant.RoomsField
	return nil
}

func ParseRooms(graph *utils.Graph, line string) error {
	if IsStart(line) && !graph.Data.StartFound {
		graph.Data.StartFound = true
	} else if IsEnd(line) && !graph.Data.EndFound {
		graph.Data.EndFound = true
	} else {
		room, err := GetRoom(graph, line)
		if err != nil {
			return err
		}
		if graph.Rooms[room] != nil {
			return &ErrorMessage{Msg: constant.ErrPrefix + "the room " + room + " is dupplicated"}
		}
		if room != "" {
			node := &utils.Room{Name: room, Neighbors: make([]*utils.Room, 0)}
			if graph.Data.StartFound && graph.Col.StartingRoom == "" {
				graph.Col.StartingRoom = room
			} else if graph.Data.EndFound && graph.Col.EndingRoom == "" {
				graph.Col.EndingRoom = room
			}
			graph.Rooms[room] = node
			graph.Col.Rooms = append(graph.Col.Rooms, node.Name)
		} else if graph.Col.StartingRoom != "" && graph.Col.EndingRoom != "" {
			graph.Data.Phase = constant.LinksField
			graph.Data.Coords = nil // free up memory from rooms coords because they are unusable
			return ParseLine(graph, line)
		} else {
			return &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrNoStart + " or " + constant.ErrNoEnd + " or " + constant.ErrSpace}
		}
	}
	return nil
}

func ParseLinks(graph *utils.Graph, line string) error {
	if graph.Col.Tunnels == nil {
		graph.Col.Tunnels = make(map[string][]string)
	}
	firstRoom, secondRoom := GetLink(line)
	if firstRoom == "" || secondRoom == "" {
		return &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrLink}
	}
	if firstRoom == secondRoom {
		return &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrLink}
	}
	if graph.Rooms[firstRoom] == nil || graph.Rooms[secondRoom] == nil {
		return &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrLink}
	}
	node1 := graph.Rooms[firstRoom]
	node2 := graph.Rooms[secondRoom]

	if DupplicatedLink(secondRoom, graph.Col.Tunnels[firstRoom]) || DupplicatedLink(firstRoom, graph.Col.Tunnels[secondRoom]) {
		return &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrLink + " at Room " + firstRoom + " , " + secondRoom}
	}
	graph.Col.Tunnels[firstRoom] = append(graph.Col.Tunnels[firstRoom], secondRoom)
	graph.Col.Tunnels[secondRoom] = append(graph.Col.Tunnels[secondRoom], firstRoom)

	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)

	// free
	node1 = nil
	node2 = nil

	return nil
}
