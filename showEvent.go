package main

import "fmt"

func showEvents() {

	fmt.Println("Voulez-vous voir tous vos évènements ou un en particulier ? (Tous/Un)")
	var choice string
	fmt.Scan(&choice)

	if choice == "Tous" {
		fmt.Println("Voici la liste des évènements : \n")
		for i := 1; i < idEvent; i++ {
			fmt.Println("------------------------------------------------\n")
			fmt.Println("Evènement n°", i)
			fmt.Println("Titre : ", eventsMap[i].title)
			fmt.Println("Date : ", eventsMap[i].date)
			fmt.Println("Heure : ", eventsMap[i].hour)
			fmt.Println("Lieu : ", eventsMap[i].place)
			fmt.Println("Catégorie : ", eventsMap[i].category)
			fmt.Println("Description : ", eventsMap[i].description)
			fmt.Println("------------------------------------------------\n")
		}
	} else if choice == "Un" {
		fmt.Println("Entrez le numéro de l'évènement : ")
		var id int
		fmt.Scan(&id)
		fmt.Println("------------------------------------------------\n")
		fmt.Println("Evènement n°", id)
		fmt.Println("Titre : ", eventsMap[id].title)
		fmt.Println("Date : ", eventsMap[id].date)
		fmt.Println("Heure : ", eventsMap[id].hour)
		fmt.Println("Lieu : ", eventsMap[id].place)
		fmt.Println("Catégorie : ", eventsMap[id].category)
		fmt.Println("Description : ", eventsMap[id].description)
		fmt.Println("------------------------------------------------\n")
	} else {
		fmt.Println("Choix invalide")
		showEvents()
	}
	menu()
}
