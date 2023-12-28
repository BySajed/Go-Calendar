package main

import (
	"fmt"
	"time"
)

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
