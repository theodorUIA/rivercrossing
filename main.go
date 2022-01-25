package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/theodoruia/rivercrossing/addremove"

	"github.com/theodoruia/myqoute"
	"github.com/theodoruia/rivercrossing/state"
)

//Globale variabler
var chicken = "kylling"
var fox = "rev"
var grain = "korn"
var hs = "mann"
var left_shore = []string{chicken, fox, grain, hs}
var right_shore = []string{}
var boat = []string{}

func ChangeState(userinput string) {
	var numchar = len(strings.Fields(userinput))

	if numchar != 5 {
		fmt.Println("Du må skrive inn 5 ord. Skriv 'hjelp' for hjelp.")
		return
	}
	//setter brukerinput til små bokstaver
	userinput = strings.ToLower(userinput)

	if userinput == "kryss" {
		state.CheckBoatContent(boat)
	}

	//split userinput
	split := strings.Split(userinput, " ")

	var item = split[0]
	var source = split[2]
	var target = split[4]
	var currentState = source + " " + target

	//Bruker skriver: Kylling fra venstre til båt
	//Item = kylling
	//Source = venstre
	//Target = båt

	switch currentState {
	case "venstre båt": //source target
		if addremove.Contains(left_shore, item) {
			left_shore = addremove.Remove(left_shore, item) //Fjerner item fra source
			boat = addremove.Add(boat, item)                //Legger item til target
			fmt.Println(item + "er fjernet fra " + source + " og legges til " + target)
			fmt.Println(state.PrintState(left_shore, right_shore, boat))

		} else {
			fmt.Println(item + " er ikke til " + source + "..")
		}
	case "båt venstre": //source target
		if addremove.Contains(boat, item) {
			boat = addremove.Remove(boat, item)          //Fjerner item fra source
			left_shore = addremove.Add(left_shore, item) //Legger item til target
			fmt.Println(item + "er fjernet fra " + source + " og legges til " + target)
			fmt.Println(state.PrintState(left_shore, right_shore, boat))

		} else {
			fmt.Println(item + " er ikke til " + source + "..")
		}
	case "båt høyre": //source target
		if addremove.Contains(boat, item) {
			boat = addremove.Remove(boat, item)            //Fjerner item fra source
			right_shore = addremove.Add(right_shore, item) //Legger item til target
			fmt.Println(item + "er fjernet fra " + source + " og legges til " + target)
			fmt.Println(state.PrintState(left_shore, right_shore, boat))

		} else {
			fmt.Println(item + " er ikke til " + source + "..")
		}
	case "høyre båt": //source target
		if addremove.Contains(right_shore, item) {
			right_shore = addremove.Remove(right_shore, item) //Fjerner item fra source
			boat = addremove.Add(boat, item)                  //Legger item til target
			fmt.Println(item + "er fjernet fra " + source + " og legges til " + target)
			fmt.Println(state.PrintState(left_shore, right_shore, boat))

		} else {
			fmt.Println(item + " er ikke til " + source + "..")
		}
	default:
		fmt.Println("Du kan skrive 'Kylling fra venstre til båt' eller 'Kylling fra båt til venstre'")
	}
}

func main() {
	var helpmsg = ("#######################################################################################################\n#                                VELKOMMEN TIL RIVERCROSSING!                                         #\n#                                                                                                     #\n# Målet med spillet er å flytte alle fra venstre til høyre med kommandoer.                            #\n#                                                                                                     #\n# Du kan skrive f.eks 'Kylling fra venstre til båt' for å flytte kylling fra venstre side oppi båten. #\n#                                                                                                     #\n# Du kan ikke flytte ting fra venstre til høyre uten å være innom båten.                              #\n# Så kommandoen 'Mann fra venstre til høyre' vil være ugyldig.                                        #\n#                                                                                                     #\n# Du kan alltid skrive 'hjelp' for hjelp eller 'state' for nåværende tilstand.                         #\n#                                                                                                     #\n#                          Dagens sitat: '" + myqoute.GetQuote() + "'                    #\n#######################################################################################################")
	fmt.Println(helpmsg)
	fmt.Println(state.PrintState(left_shore, right_shore, boat))

	userinput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nHva vil du gjøre? \n") // reads user input until \n by default
		userinput.Scan()                    // Holds the string that was scanned
		text := userinput.Text()
		if text == "hjelp" {
			fmt.Println(helpmsg)
		} else if text == "state" {
			state.PrintState(left_shore, right_shore, boat)
		} else {
			if len(text) != 0 {
				//fmt.Println(text)
				ChangeState(text)
			} else {
				fmt.Println("Du må skrive noe! Du kan skrive 'Kylling fra venstre til båt' eller 'Kylling fra båt til venstre'. Skriv 'hjelp' for hjelp.")
			}
		}

	}
}
