package state

import (
	"fmt"
	"strings"
)

func PrintState(left []string, right []string, boat []string) string {
	left_string := strings.Join(left, ", ")
	right_string := strings.Join(right, ", ")
	boat_string := strings.Join(boat, "")

	var statemsg = ("#########################################################################################################\n#                                                                                                     #\n#  Venstre: " + left_string + "                   Båt: \"" + boat_string + "/       Høyre: " + right_string + "\n#                                                                                                     #\n#######################################################################################################")

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
