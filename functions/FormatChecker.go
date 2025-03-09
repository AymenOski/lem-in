package functions

import (
	"fmt"
	"os"
)

func FormatChekcer(FC_AntNumber int, FC_StartFlag int, FC_EndFlag int, EOF_line int) {
	switch {
	case FC_AntNumber == -1: //-----------AntNum-------------
		fmt.Println("You misplaced the ant number in the file.")
		os.Exit(0)
	case FC_StartFlag > FC_EndFlag: //----------##end & ##start-------------
		fmt.Println("ERROR: invalid data format, '##start' appears after '##end'.")
		os.Exit(0)
	case FC_StartFlag == 0 ||
		FC_StartFlag == EOF_line: //--------------##start-----------
		fmt.Println("ERROR: invalid data format, Missing or misplaced '##start' in the file.")
		os.Exit(0)
	case FC_EndFlag == 0 ||
		FC_EndFlag == EOF_line: //---------------##end----------
		fmt.Println("ERROR: invalid data format, Missing or misplaced '##end' in the file.")
		os.Exit(0)
	}
}
