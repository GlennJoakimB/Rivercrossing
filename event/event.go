package event

import (
	"fmt"

	"github.com/GlennJoakimB/rivercrossing/state"
)

//PutInBoat
func MoveItem(item string) {
	ky, re, ko, ma, bo := state.GetViewState()
	bItem := state.GetBoatItem()

	if item == "Mann" {
		if ma != 1 {
			// hvis mann ikke er i båt, putt i båt
			ma = 1
		} else {
			// ta man ut av båt
			ma = bo
		}
	} else if bItem == "_" {
		// hvis båt er empty (utenom mann)
		// putt item i båt
		//men bare hvis båten er der for å plukke dem opp!
		if item == "Kylling" {
			if ky == bo { //er kylling og båt på samme side?
				ky = 1
			} else {
				tooFar(item)
				return
			}
		} else if item == "Rev" {
			if re == bo {
				re = 1
			} else {
				tooFar(item)
				return
			}
		} else if item == "Korn" {
			if re == bo {
				ko = 1
			} else {
				tooFar(item)
				return
			}
		} else {
			fmt.Println(item, "er ikke gjenkjent. Prøv igjen.")
			return
		}
	} else if bItem == item {
		// Hvis item allerede er i båt
		// ta ut av båt og putt på land der båten står
		if item == "Kylling" {
			ky = bo
		}
		if item == "Rev" {
			re = bo
		}
		if item == "Korn" {
			ko = bo
		}
	} else {
		// hvis flytting av item ikke kan skje
		if ma != 1 {
			// hvis mann ikke er med
			fmt.Printf("%v kan ikke ro båten alene. Mannen må være med.\n", bItem)
		} else {
			fmt.Println("Båten har bare plass til to om gangen.")
		}
		return
	}
	// Oppdater viewstate
	state.ViewState(ky, re, ko, ma, bo)
}

//Funksjon for å flytte båt
func Cross() string {
	ky, re, ko, ma, bo := state.GetViewState()

	if ma == 1 {
		if bo == 0 {
			bo = 2
		} else { //if bo == 2 "Øst"
			bo = 0
		}
	} else {
		fmt.Println("Mannen må være med for å styre båten")
		return "nullMann"
	}
	state.ViewState(ky, re, ko, ma, bo)

	// Sjekk om spiller har tapt når båt seiler
	if ky == re && ky == ko && ma == 1 {
		// Hvis mannen har seilt alene
		fmt.Println("Ops! Kornet ble spist av kyllingen, og kyllingen ble spist av reven!")
		return "GameOver"
	} else if ky == re && ma == 1 {
		// Hvis kylling og rev står på samme sted, og uten mann tilstede
		fmt.Println("Ops! Reven spiser kyllingen!")
		return "GameOver"
	} else if ky == ko && ma == 1 {
		// Hvis kylling er alene med korn, uten mannen til stede
		fmt.Println("Ops! Kyllingen spiser kornet!")
		return "GameOver"
	} else {
		return "noError"
	}
}

// Print hvis item er på feil side relativt til båten.
func tooFar(item string) {
	fmt.Println(item, "er dessverre på feil side. Prøv noe annet.")
}
