package main

import (
	"fmt"
	"os"
)

var input string

func main() {
	fmt.Println("Hallo")
	ask()
}

func ask() {
	dict := map[string]string{
		"name":  "chatbot",
		"alter": "ich wurde 22.02.26 erstellt",
		"warum": "als miniprojekt zum Go lernen",
		"aus":   "Programm wird beendet",
		"hilfe": "info		-> informationen\naus		-> beenden\n",
		"info":  "Mögliche Commands:\nname - gibt den namen aus\nwarum - Grund, wieso ich diesen chatbot erstellte\naus - beenden\n",
	}

	for i := 0; i < 9999; {

		fmt.Scanf("%s", &input)
		_, vorhanden := dict[input]

		if input == "aus" {
			turn_off()
		} else if vorhanden {
			fmt.Printf("%s", dict[input])
		} else {
			fmt.Println("Kein Passender Eintrag gefunden")
		}
	}
}

func turn_off() {
	fmt.Println("Programm wird beendet")
	os.Exit(0)
}
