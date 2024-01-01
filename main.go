package main

import (
	"fmt"
)

type Event struct {
	title       string
	date        string
	hour        string
	place       string
	category    string
	description string
}

var eventsMap = make(map[int]Event)
var idEvent = 1

func menu() {
	choice := 0
	cExit := 0

	fmt.Println("\n------------------------------------------------\n")
	fmt.Println("1. Créer un nouvel évènement\n")
	fmt.Println("2. Visualiser les évènements\n")
	fmt.Println("3. Modifier un évènement\n")
	fmt.Println("4. Supprimer un évènement\n")
	fmt.Println("5. Rechercher un évènement\n")
	fmt.Println("6. Quitter\n")
	fmt.Println("Chosissez une option : ")

	for cExit != 1 {
		fmt.Scan(&choice)
		switch choice {

		case 1:
			newEvent()
			choice = 0
			menu()

		case 2:
			showEvents()
			choice = 0
			menu()

		case 3:
			modifyEvent()
			choice = 0
			menu()

		case 4:
			/* TODO: Supprimer un évènement */
			choice = 0
			menu()

		case 5:
			/* TODO: Recherchez un évènment */
			choice = 0
			menu()

		case 6:
			/* TODO: Quitter */
			choice = 0
			cExit++
		default:
			fmt.Println("Entrée invalide, veuillez réessayez")
		}
	}
}

func main() {
	fmt.Println("Bienvenue dans le système de gestion de planning\n")
	menu()
}
