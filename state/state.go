package state

import (
	"strings"
)

func PrintState(left []string, right []string, boat []string) string {
	left_string := strings.Join(left, ", ")
	right_string := strings.Join(right, ", ")
	boat_string := strings.Join(boat, "")

	var statemsg = ("\nNåværende tilstand er: \n#########################################################################################################\n#                                                                                                     #\n#  Venstre: " + left_string + "                   Båt: \"" + boat_string + "/       Høyre: " + right_string + "\n#                                                                                                     #\n#######################################################################################################")

	return statemsg
}

func IsBoatEmpty(boat []string) bool {
	if len(boat) == 0 {
		return true
	}
	return false
}
