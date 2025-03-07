package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExtractingDate() (int, int, int, map[int][]int) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please create a file in the root of the project and put input to it")
		os.Exit(0)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	c := 0
	antNum := 0
	startingRoom := 0
	endingRoom := 0
	Tunnels := make(map[int][]int)
	check := false
	for scanner.Scan() {
		if c == 0 {
			antNum, _ = strconv.Atoi(scanner.Text())
			c++
		}
		if c == 3 {
			// storing data after the ##end flag
			if len(scanner.Text()) < 4 {
				temp, _ := strconv.Atoi(scanner.Text()[:1])
				temp2, _ := strconv.Atoi(scanner.Text()[2:3])
				Tunnels[temp] = append(Tunnels[temp], temp2)
			}
		}
		if check {
			startingRoom, _ = strconv.Atoi(scanner.Text()[:1])
			check = false
		} else if !check && c == 2 {
			endingRoom, _ = strconv.Atoi(scanner.Text()[:1])
			c = 3
		}
		if strings.HasPrefix(scanner.Text(), "##start") {
			check = true
		} else if strings.HasPrefix(scanner.Text(), "##end") {
			c = 2
		}
	}
	return antNum, startingRoom, endingRoom, Tunnels
}
