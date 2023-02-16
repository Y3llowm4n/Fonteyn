package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	kenteken string
)

// hierin wordt de input aangemaakt, waar de gebruiker zijn/haar kenteken kan invullen.
func init() {
	flag.StringVar(&kenteken, "k", "MG-21-YP", "Voorbeeld: MG-21-YP")
	flag.Parse()
}

func main() {
	greet := realTime()
	check(greet)
}

// hierin wordt gecheckt of het kenteken in de lijst staat.
func check(greet int) {
	kentekens := []string{"MG-21-YP", "YP-25-MG"}

	found := false
	for _, number := range kentekens {
		if number == kenteken {
			found = true
			break
		}
	}

	//Als het kenteken matched wordt er een welkomsbericht gegeven. Zo niet dan heeft de persoon geen toegang.
	if found {
		fmt.Println(greeting(greet))
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein.")
	}

}

// geeft de realistische tijd weer in uren
func realTime() int {
	currentTime := time.Now()
	groet := currentTime.Hour()
	return groet
}

func greeting(moment int) string {
	switch {
	case moment >= 7 && moment <= 12:
		return "Goedenmorgen"
	case moment >= 12 && moment <= 18:
		return "Goedenmiddag"
	case moment >= 18 && moment <= 24:
		return "Goedenavond"
	default:
		return "Sorry, de parkeerplaats	is 's nachts gesloten"
	}

}
