//  Het bericht krijgt het formaat: [groet]! Welkom bij Fonteyn Vakantieparken
// Waarbij [groet] een plaataanduiding is voor
// Goedemorgen als de huidige tijd tussen 7 uur ’s ochtends en 12 uur ’s middags ligt
// Goedemiddag als de huidige tijd tussen 12 uur ’s middags en 18 uur ’s middags ligt
// Goedenavond als de huidige tijd tussen 18 uur ’s middags en 23 uur ’s avonds ligt
// Als het tussen 23 uur ’s avonds en 7 uur ’s morgens is, dient het bericht te zijn: Sorry, de parkeerplaats is ’s nachts gesloten
package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	groet := currentTime.Hour()

	fmt.Println(groet)
	if groet >= 7 && groet <= 11 {
		fmt.Println("Goedenmorgen")
	} else if groet >= 12 && groet <= 18 {
		fmt.Println("Goedenmiddag")
	} else if groet >= 18 && groet <= 23 {
		fmt.Println("Goedenavond!")
	} else {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")
	}
}
