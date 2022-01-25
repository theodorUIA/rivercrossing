package state

import (
	"fmt"
)

func PrintState(left []string, right []string, boat []string) string {
	fmt.Println("\nNåværende tilstand:")
	fmt.Println("Til venstre er: ", left)
	fmt.Println("Og i båten er: ", boat)
	fmt.Println("Til høyre er: ", right)
	return ""
}

func CheckBoatContent(boat []string) string {
	if len(boat) == 0 {
		return "båten er tom"
	} else {
		return "båten er ikke tom"
	}
}
