package constant

const (
	ErrPrefix    = "ERROR: invalid data format, "
	ErrArgs      = "invalid number of arguments"
	ErrFileIssue = "couldn't open the file"
	ErrData      = "invalid data format."
	ErrRoomName  = "room name shouldn't start with L and be empty"
	ErrAnts      = "invalid number of ants"
	ErrCoord     = "invalid coordinates"
	ErrNoPaths   = "no valid paths were found"
	ErrNoStart   = "no start room found"
	ErrNoEnd     = "no end room found"
	ErrLink      = "invalid link"
	ErrSpace     = "a room or more have spaces"
)

const (
	AntsField  = iota // 0
	RoomsField        // 1
	LinksField        // 2
)
