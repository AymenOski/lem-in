package functions

import (
	"fmt"
	"os"
)

func FormatChekcer(FC_AntNumber int, FC_StartFlag int, FC_EndFlag int, Rooms []string, EOF_line int) {
	switch {
	case FC_StartFlag == 0 ||
		FC_StartFlag == EOF_line: //--------------##start-----------
		fmt.Println("ERROR: invalid data format, Missing or misplaced '##start' in the file.")
		os.Exit(0)
	case FC_EndFlag == 0 ||
		FC_EndFlag == EOF_line: //---------------##end----------
		fmt.Println("ERROR: invalid data format, Missing or misplaced '##end' in the file.")
		os.Exit(0)
	case FC_StartFlag > FC_EndFlag: //----------##end & ##start-------------
		fmt.Println("ERROR: invalid data format, '##start' appears after '##end'.")
		os.Exit(0)
	}
	//---------------duplicated Rooms----------
	for i, room1 := range Rooms {
		for j, room2 := range Rooms {
			if i != j && room1 == room2 {
				fmt.Println("ERROR: duplicate room name in the file.")
				os.Exit(0)
			}
		}
	}
	//---------------duplicated Coordinates----------
}
