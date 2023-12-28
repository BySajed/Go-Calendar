package main

import "fmt"

func modifyEvent() {
	fmt.Println("Entrez le numéro de l'évènement que vous souhaitez modifier : ")
	var id int
	fmt.Scan(&id)
	event := eventsMap[id]
	fmt.Println("------------------------------------------------\n")
	fmt.Println("Evènement n°", id)
	fmt.Println("Titre : ", eventsMap[id].title)
	fmt.Println("Date : ", eventsMap[id].date)
	fmt.Println("Heure : ", eventsMap[id].hour)
	fmt.Println("Lieu : ", eventsMap[id].place)
	fmt.Println("Catégorie : ", eventsMap[id].category)
	fmt.Println("Description : ", eventsMap[id].description)
	fmt.Println("------------------------------------------------\n")

	fmt.Println("Que souhaitez-vous modifier ? (Titre/Date/Heure/Lieu/Catégorie/Description)")
	var choice string
	fmt.Scan(&choice)

	if choice == "Titre" || choice == "titre" {
		fmt.Println("Entrez le nouveau titre : ")
		var newTitle string
		fmt.Scan(&newTitle)
		event.title = newTitle
		fmt.Println("Titre modifié avec succès !")

	} else if choice == "Date" || choice == "date" {
		fmt.Println("Entrez la nouvelle date : ")
		var newDate string
		fmt.Scan(&newDate)
		event.date = newDate
		fmt.Println("Date modifiée avec succès !")

	} else if choice == "Heure" || choice == "heure" {
		fmt.Println("Entrez la nouvelle heure : ")
		var newHour string
		fmt.Scan(&newHour)
		event.hour = newHour
		fmt.Println("Heure modifiée avec succès !")

	} else if choice == "Lieu" || choice == "lieu" {
		fmt.Println("Entrez le nouveau lieu : ")
		var newPlace string
		fmt.Scan(&newPlace)
		event.place = newPlace
		fmt.Println("Lieu modifié avec succès !")

	} else if choice == "Catégorie" || choice == "catégorie" {
		fmt.Println("Entrez la nouvelle catégorie : ")
		var newCategory string
		fmt.Scan(&newCategory)
		event.category = newCategory
		fmt.Println("Catégorie modifiée avec succès !")

	} else if choice == "Description" || choice == "description" {
		fmt.Println("Entrez la nouvelle description : ")
		var newDescription string
		fmt.Scan(&newDescription)
		event.description = newDescription
		fmt.Println("Description modifiée avec succès !")
	} else {
		fmt.Println("Choix invalide")
		modifyEvent()
	}

	eventsMap[id] = event
}
