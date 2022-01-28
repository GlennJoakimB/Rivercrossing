package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/GlennJoakimB/rivercrossing/event"
	"github.com/GlennJoakimB/rivercrossing/state"
)

func main() {
	// state.ViewState(gameState)
	// fmt.Println(state.ViewState())
	initState()
}

func initState() {
	fmt.Println("") //mellomrom
	state.ViewState(0, 0, 0, 0, 0)
	GetInput()
}

func GetInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("") // Tomt rom
	// fmt.Println("(For tips, skriv Hjelp)")
	fmt.Print("Hva vil du gjøre? ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	response(input)
}

func response(input string) {
	gState := ""
	if input == "Stopp" || input == "Stop" {
		return
	} else if input == "Hjelp" {
		fmt.Println(" Stopp | stopper spillet.")
		fmt.Println(" Restart | starter på nytt.")
		fmt.Println(" Status | viser state til spillet.")
		fmt.Println(" Kylling | vil flytte kylling inn/ut av båten.")
		fmt.Println(" Kryss elv | vil sende båten over til andre siden.")
	} else if input == "Restart" {
		initState()
		return
	} else if input == "Status" {
		ky, re, ko, ma, bo := state.GetViewState()
		state.ViewState(ky, re, ko, ma, bo)
	} else if input == "Kryss elv" {
		gState = event.Cross()
	} else {
		event.MoveItem(input)
	}

	win := state.GetWin()
	if win == "Win" {
		fmt.Println("")
		fmt.Println("Gratulerer!")
		fmt.Println("Alle fire kom seg over!")
	} else if gState == "GameOver" {
		fmt.Println("GameOver")
		fmt.Println("")
	} else {
		GetInput()
	}
}
