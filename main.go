package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Globale variabler
var chicken = "kylling"
var fox = "rev"
var grain = "korn"
var hs = "mann"
var left_shore = []string{chicken, fox, grain, hs}
var right_shore = []string{}
var boat = []string{}

func ViewState() string {
	fmt.Println("\nNåværende tilstand:")
	fmt.Println("Til venstre er: ", left_shore)
	fmt.Println("Og i båten er: ", boat)
	fmt.Println("Til høyre er: ", right_shore)
	return ""
}

func ChangeState(userinput string) {
	var numchar = len(strings.Fields(userinput))

	if numchar != 5 {
		fmt.Println("Du må skrive inn 5 ord. Skriv 'help' for hjelp.")
		return
	}

	userinput = strings.ToLower(userinput)
	//split userinput
	split := strings.Split(userinput, " ")

	var item = split[0]
	var source = split[2]
	var target = split[4]
	var state = source + " " + target

	//Bruker skriver: Kylling fra venstre til båt
	//Item = kylling
	//Source = venstre
	//Target = båt

	switch state {
	case "venstre båt": //source target
		if Contains(left_shore, item) {
			left_shore = Remove(left_shore, item) //Fjerner item fra source
			boat = Add(boat, item)                //Legger item til target
			fmt.Println(item + "er fjernet fra " + source + " og legges til " + target)
			fmt.Println(ViewState())

		} else {
			fmt.Println(item + " er ikke til " + source + "..")
		}
	case "båt venstre": //source target
		if Contains(boat, item) {
			boat = Remove(boat, item)          //Fjerner item fra source
			left_shore = Add(left_shore, item) //Legger item til target
			fmt.Println(item + "er fjernet fra " + source + " og legges til " + target)
			fmt.Println(ViewState())

		} else {
			fmt.Println(item + " er ikke til " + source + "..")
		}
	case "båt høyre": //source target
		if Contains(boat, item) {
			boat = Remove(boat, item)            //Fjerner item fra source
			right_shore = Add(right_shore, item) //Legger item til target
			fmt.Println(item + "er fjernet fra " + source + " og legges til " + target)
			fmt.Println(ViewState())

		} else {
			fmt.Println(item + " er ikke til " + source + "..")
		}
	case "høyre båt": //source target
		if Contains(right_shore, item) {
			right_shore = Remove(right_shore, item) //Fjerner item fra source
			boat = Add(boat, item)                  //Legger item til target
			fmt.Println(item + "er fjernet fra " + source + " og legges til " + target)
			fmt.Println(ViewState())

		} else {
			fmt.Println(item + " er ikke til " + source + "..")
		}
	default:
		fmt.Println("Du kan skrive 'Kylling fra venstre til båt' eller 'Kylling fra båt til venstre'")
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func Add(s []string, r string) []string {
	s = append(s, r)
	return s
}

func main() {
	fmt.Println("########################################################### \n# Målet med spillet er å flytte alle fra venstre til høyre med kommandoer. \n#\n# Du kan skrive f.eks 'Kylling fra venstre til båt' for å flytte kylling fra venstre side oppi båten. \n# \n# Du kan ikke flytte ting fra venstre til høyre uten å være innom båten. Så kommandoen 'Mann fra venstre til høyre' vil være ugyldig. \n#\n# Du kan alltid skrive 'help' for hjelp eller 'state' for nåværende tilstand. \n###########################################################")
	fmt.Println(ViewState())

	userinput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nHva vil du gjøre? \n") // reads user input until \n by default
		userinput.Scan()                    // Holds the string that was scanned
		text := userinput.Text()
		if text == "help" {
			fmt.Println("########################################################### \n# Målet med spillet er å flytte alle fra venstre til høyre med kommandoer. \n#\n# Du kan skrive f.eks 'Kylling fra venstre til båt' for å flytte kylling fra venstre side oppi båten. \n# \n# Du kan ikke flytte ting fra venstre til høyre uten å være innom båten. Så kommandoen 'Mann fra venstre til høyre' vil være ugyldig. \n#\n# Du kan alltid skrive 'help' for hjelp eller 'state' for nåværende tilstand. \n###########################################################")
		} else if text == "state" {
			ViewState()
		} else {
			if len(text) != 0 {
				//fmt.Println(text)
				ChangeState(text)
			} else {
				fmt.Println("Du må skrive noe! Du kan skrive 'Kylling fra venstre til båt' eller 'Kylling fra båt til venstre'. Skriv 'help' for hjelp.")
			}
		}

	}
}
