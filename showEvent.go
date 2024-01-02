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
			fmt.Println("Titre : ", eventsMap[i].Title)
			fmt.Println("Date : ", eventsMap[i].Date)
			fmt.Println("Heure : ", eventsMap[i].Hour)
			fmt.Println("Lieu : ", eventsMap[i].Place)
			fmt.Println("Catégorie : ", eventsMap[i].Category)
			fmt.Println("Description : ", eventsMap[i].Description)
		}
	} else if choice == "Un" {
		fmt.Println("Entrez le numéro de l'évènement : ")
		var id int
		fmt.Scan(&id)
		fmt.Println("------------------------------------------------\n")
		fmt.Println("Evènement n°", id)
		fmt.Println("Titre : ", eventsMap[id].Title)
		fmt.Println("Date : ", eventsMap[id].Date)
		fmt.Println("Heure : ", eventsMap[id].Hour)
		fmt.Println("Lieu : ", eventsMap[id].Place)
		fmt.Println("Catégorie : ", eventsMap[id].Category)
		fmt.Println("Description : ", eventsMap[id].Description)
	} else {
		fmt.Println("Choix invalide")
		showEvents()
	}
	//menu()
	return
}
