package state

import (
	"fmt"
	"strings"
)

func PrintState(left []string, right []string, boat []string) string {
	left_string := strings.Join(left, ", ")
	right_string := strings.Join(right, ", ")
	boat_string := strings.Join(boat, "")

	var statemsg = ("#######################################################################################################\n# _______________                                                                     _______________ #\n# | Venstre:     |            \"   Båt:           /                                   |               #\n#  " + left_string + "               " + boat_string + "                                      " + right_string + "#\n# |______________|               \"_____________/                                     |               #\n#######################################################################################################")

	fmt.Println("\nNåværende tilstand:")

	fmt.Println(statemsg)
	return ""
}

func IsBoatEmpty(boat []string) bool {
	if len(boat) == 0 {
		return true
	}
	return false
}
