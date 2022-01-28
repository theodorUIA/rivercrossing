package state

import (
	"strings"

	"github.com/theodoruia/rivercrossing/addremove"
)

func PrintState(left []string, right []string, boat []string) string {
	left_string := strings.Join(left, ", ")
	right_string := strings.Join(right, ", ")
	boat_string := strings.Join(boat, ", ")

	var statemsg = ("\nNåværende tilstand er: \n#######################################################################################################\n#                                                                                                     #\n#  Venstre: " + left_string + "                   Båt: " + boat_string + "                Høyre: " + right_string + "\n#                                                                                                     #\n#######################################################################################################")

	return statemsg
}

//Sjekker om båten allerede har 2 elementer i seg
func IsBoatFull(boat []string) bool {
	if len(boat) <= 1 {
		return true
	}
	return false
}

func CheckWin(left_shore []string, boat []string) bool {
	if len(left_shore) == 0 && len(boat) == 0 {
		return true
	}
	return false
}
func CheckLose(left_shore []string, boat []string, right_shore []string, chicken string, fox string, grain string) bool {
	//Sjekker om venstre eller båt har to "items" i seg
	if addremove.Contains(left_shore, chicken) && addremove.Contains(left_shore, fox) {
		return true
	}
	if addremove.Contains(boat, chicken) && addremove.Contains(boat, fox) {
		return true
	}
	if addremove.Contains(right_shore, chicken) && addremove.Contains(right_shore, fox) {
		return true
	}
	if addremove.Contains(left_shore, chicken) && addremove.Contains(left_shore, grain) {
		return true
	}
	if addremove.Contains(boat, chicken) && addremove.Contains(boat, grain) {
		return true
	}
	if addremove.Contains(right_shore, chicken) && addremove.Contains(right_shore, grain) {
		return true
	}
	return false
}
