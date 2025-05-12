package parse

import (
	"strconv"
	"strings"

	constant "lem-in/const"
	"lem-in/utils"
)

func IsComment(s string) bool {
	return strings.HasPrefix(s, "#") && !IsStart(s) && !IsEnd(s)
}

func IsStart(s string) bool {
	return s == "##start"
}

func IsEnd(s string) bool {
	return s == "##end"
}

func IsRoom(s string) bool {
	return strings.HasPrefix(s, "L")
}

func IsCoords(x, y string, graph *utils.Graph) bool {
	xn, err := strconv.Atoi(x)
	yn, err := strconv.Atoi(y)
	if err != nil {
		return false
	}
	if ok := graph.Data.Coords[[2]int{xn, yn}]; ok {
		return false
	}
	graph.Data.Coords[[2]int{xn, yn}] = true
	return true
}

func GetRoom(graph *utils.Graph, s string) (string, error) {
	split := strings.Split(s, " ")
	if len(split) != 3 {
		return "", nil // to dodge links info if they comes first in order
	}
	if IsRoom(split[0]) {
		return "", &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrRoomName}
	}
	if !IsCoords(split[1], split[2], graph) {
		return "", &ErrorMessage{Msg: constant.ErrPrefix + constant.ErrCoord + " : dupplicated " + split[0]}
	}
	return split[0], nil
}

func GetLink(s string) (string, string) {
	split := strings.Split(s, "-")
	if len(split) != 2 {
		return "", ""
	}
	return split[0], split[1]
}

func DupplicatedLink(s string, arr []string) bool {
	for _, link := range arr {
		if link == s {
			return true
		}
	}
	return false
}
