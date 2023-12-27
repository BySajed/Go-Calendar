package main

import (
	"fmt"
	"time"
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

	fmt.Println("Bienvenue dans le système de gestion de planning\n")
	fmt.Println("------------------------------------------------\n")
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

		case 2:
			showEvents()
			choice = 0

		case 3:
			modifyEvent()
			choice = 0

		case 4:
			/* TODO: Supprimer un évènement */
			choice = 0

		case 5:
			/* TODO: Recherchez un évènment */
			choice = 0

		case 6:
			/* TODO: Quitter */
			choice = 0
			cExit++

		default:
			fmt.Println("Choix invalide")
			menu()
		}
	}
}

func newEvent() {
	var title string
	var date string
	fmt_date := "2006-01-02"
	var hour string
	var place string
	var category string
	var description string

	fmt.Println("Entrez le titre de l'évènement : ")
	fmt.Scan(&title)
	fmt.Println("Entrez la date de l'évènement (YYYY-MM-DD): ")
	fmt.Scan(&date)

	//Gestion d'erreur en cas de mauvais format de date
	parsedTime, err := time.Parse(fmt_date, date)
	if err != nil {
		fmt.Println("Date invalide")
		newEvent()
	}
	if parsedTime.Format(fmt_date) != date {
		fmt.Println("e format date est invalide")
		newEvent()
	}

	//Gestion d'erreur en cas de date antérieure à la date du jour
	if parsedTime.Before(time.Now()) {
		fmt.Println("Date invalide")
		newEvent()
	}

	fmt.Println("Entrez l'heure de l'évènement (HH:MM): ")
	fmt.Scan(&hour)

	//Gestion d'erreur en cas de mauvais format d'heure
	parsedHour, err := time.Parse("15:04", hour)
	if err != nil {
		fmt.Println("Heure invalide")
		newEvent()
	}
	if parsedHour.Format("15:04") != hour {
		fmt.Println("Le format heure est invalide")
		newEvent()
	}

	fmt.Println("Entrez le lieu de l'évènement : ")
	fmt.Scan(&place)

	fmt.Println("Choisissez une catégorie(Professionnel, Personnel, Loisir): ")
	fmt.Scan(&category)
	if category != "Professionnel" && category != "Personnel" && category != "Loisir" {
		fmt.Println("Catégorie invalide")
		newEvent()
	}

	fmt.Println("Entrez une brève description de l'évènement : ")
	fmt.Scan(&description)

	newEvent := Event{
		title:       title,
		date:        date,
		hour:        hour,
		place:       place,
		category:    category,
		description: description,
	}
	eventsMap[idEvent] = newEvent
	idEvent++

	fmt.Println("Evènement créé avec succès !")
}

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

func main() {
	menu()
}
