package state

import "fmt"

var boatItem string = "_"
var boatMan string = "_"
var ky int
var re int
var ko int
var ma int
var bo int

//Update the visual state of the game
func ViewState(v1 int, v2 int, v3 int, v4 int, v5 int) {

	ky = v1
	re = v2
	ko = v3
	ma = v4
	bo = v5

	Vky := ""
	Vre := ""
	Vko := ""
	Vma := ""

	Øky := ""
	Øre := ""
	Øko := ""
	Øma := ""

	bItem := "_"
	bMann := "_"

	//Kylling
	if ky == 1 {
		bItem = "Kylling"
	} else if ky == 0 {
		Vky = "Kylling"
	} else {
		Øky = "Kylling"
	}

	//Reven
	if re == 1 {
		bItem = "Rev"
	} else if re == 0 {
		Vre = "Rev"
	} else {
		Øre = "Rev"
	}

	//Korn
	if ko == 1 {
		bItem = "Korn"
	} else if ko == 0 {
		Vko = "Korn"
	} else {
		Øko = "Korn"
	}

	//Mann
	if ma == 1 {
		bMann = "Mann"
	} else if ma == 0 {
		Vma = "Mann"
	} else {
		Øma = "Mann"
	}

	if bo == 0 {
		fmt.Printf("[%v %v %v %v ]---\\ \\_%v_%v_/ _________________/---[ %v %v %v %v]\n", Vky, Vre, Vko, Vma, bItem, bMann, Øky, Øre, Øko, Øma)
	} else {
		fmt.Printf("[%v %v %v %v ]---\\_________________ \\_%v_%v_/ /---[ %v %v %v %v]\n", Vky, Vre, Vko, Vma, bItem, bMann, Øky, Øre, Øko, Øma)
	}

	boatItem = bItem
	boatMan = bMann

}

func GetBoatItem() string {
	return boatItem
}
func GetBoatMan() string {
	return boatMan
}
func GetViewState() (int, int, int, int, int) {
	return ky, re, ko, ma, bo
}
func GetWin() string {
	// Sjekk om spiller har vunnet
	if ky == 2 && re == 2 && ko == 2 && ma == 2 {
		return "Win"
	} else {
		return "Nono"
	}
}
