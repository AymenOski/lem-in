package utils

type Coordinates struct {
	X int
	Y int
}

type Colony struct {
	AntNum       int
	Rooms        []string
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
	Data  ParseInfo
	Col   *Colony
}

type ParseInfo struct {
	Coords               map[[2]int]bool
	Phase                int
	StartFound, EndFound bool
}
