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
