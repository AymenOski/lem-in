package constant

const (
	ErrArgs      = "ERROR: invalid number of arguments"
	ErrFileIssue = "ERROR: couldn't open the file"
	ErrData      = "ERROR: invalid data format."
	ErrRoomName  = "room name shouldn't start with L and be empty"
	ErrAnts      = "invalid number of ants"
	ErrCoord     = "invalid coordinates"
	ErrNoPaths   = "no valid paths were found"
	ErrNoStart   = "no start room found"
	ErrNoEnd     = "no end room found"
	ErrLink      = "invalid link"
)

const (
	AntsField  = iota // 0
	RoomsField        // 1
	LinksField        // 2
)
