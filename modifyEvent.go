package main

import "fmt"

func modifyEvent() {
	fmt.Println("Entrez le numéro de l'évènement que vous souhaitez modifier : ")
	var id int
	fmt.Scan(&id)
	event := eventsMap[id]
	fmt.Println("------------------------------------------------\n")
	fmt.Println("Evènement n°", id)
	fmt.Println("Titre : ", eventsMap[id].Title)
	fmt.Println("Date : ", eventsMap[id].Date)
	fmt.Println("Heure : ", eventsMap[id].Hour)
	fmt.Println("Lieu : ", eventsMap[id].Place)
	fmt.Println("Catégorie : ", eventsMap[id].Category)
	fmt.Println("Description : ", eventsMap[id].Description)
	fmt.Println("------------------------------------------------\n")

	fmt.Println("Que souhaitez-vous modifier ? (Titre/Date/Heure/Lieu/Catégorie/Description)")
	var choice string
	fmt.Scan(&choice)

	if choice == "Titre" || choice == "titre" {
		fmt.Println("Entrez le nouveau titre : ")
		var newTitle string
		fmt.Scan(&newTitle)
		event.Title = newTitle
		fmt.Println("Titre modifié avec succès !")

	} else if choice == "Date" || choice == "date" {
		fmt.Println("Entrez la nouvelle date : ")
		var newDate string
		fmt.Scan(&newDate)
		event.Date = newDate
		fmt.Println("Date modifiée avec succès !")

	} else if choice == "Heure" || choice == "heure" {
		fmt.Println("Entrez la nouvelle heure : ")
		var newHour string
		fmt.Scan(&newHour)
		event.Hour = newHour
		fmt.Println("Heure modifiée avec succès !")

	} else if choice == "Lieu" || choice == "lieu" {
		fmt.Println("Entrez le nouveau lieu : ")
		var newPlace string
		fmt.Scan(&newPlace)
		event.Place = newPlace
		fmt.Println("Lieu modifié avec succès !")

	} else if choice == "Catégorie" || choice == "catégorie" {
		fmt.Println("Entrez la nouvelle catégorie : ")
		var newCategory string
		fmt.Scan(&newCategory)
		event.Category = newCategory
		fmt.Println("Catégorie modifiée avec succès !")

	} else if choice == "Description" {
		fmt.Println("Entrez la nouvelle Description : ")
		var newDescription string
		fmt.Scan(&newDescription)
		event.Description = newDescription
		fmt.Println("Description modifiée avec succès !")
	} else {
		fmt.Println("Choix invalide")
		modifyEvent()
	}

	eventsMap[id] = event
}
