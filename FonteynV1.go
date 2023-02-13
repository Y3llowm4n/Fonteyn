// Input (het kenteken) aan het programma kan worden gegeven
// Het programma kijkt of het kenteken voorkomt in een hard-coded reeks van kentekens
// Als het kenteken voorkomt, het welkomstbericht wordt getoond zoals dat in v2 van de
// applicatie is gedefinieerd
// Als het kenteken niet voorkomt in deze reeks, de volgende tekst wordt getoond: U heeft
// helaas geen toegang tot het parkeerterrein
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
	flag.StringVar(&kenteken, "k", "MG-21-YP", "kenteken van nummerplaat")
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

	//Als het kenteken matched wordt er een welkoms bericht gegeven. Zo niet dan heeft de persoon geen toegang.
	if found {
		greeting(greet)
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

func greeting(groet int) {

	if groet >= 7 && groet <= 11 {
		fmt.Println("Goedenmorgen, Welkom bij Fonteyn Vakantieparken!")
	} else if groet >= 12 && groet <= 18 {
		fmt.Println("Goedenmiddag, Welkom bij Fonteyn Vakantieparken!")
	} else if groet >= 18 && groet <= 23 {
		fmt.Println("Goedenavond, Welkom bij Fonteyn Vakantieparken!")
	} else {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")
	}
}
