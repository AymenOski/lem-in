package utils

type Coordinates struct {
	X int
	Y int
}

type Ants struct {
	AntNum       int
	Rooms        []string
	// Position or should we change Rooms to a map of coordinates to room names?
	Position     []Coordinates
	Tunnels      map[string][]string
	StartingRoom string
	EndingRoom   string
}

// The rooms names will not necessarily be numbers, and in order.
// The rooms are identified by a string. it could be "A", "B", "1", "2", etc.  <--- rooms are not necessarily numbers
type Room struct {
	RoomName string
	AntExist bool
}

// A tunnel joins only two rooms together never more than that.
// meaning that a pointer can point to only one room at a time.
type Tunnels struct {
	Room1 *Room
	Room2 *Room
}
