package main

import (
	"fmt"
	"time"
)

func menu() {
	choice := 0

	fmt.Println("Bienvenue dans le système de gestion de planning\n")
	fmt.Println("------------------------------------------------\n")
	fmt.Println("1. Créer un nouvel évènement\n")
	fmt.Println("2. Visualiser les évènements\n")
	fmt.Println("3. Modifier un évènement\n")
	fmt.Println("4. Supprimer un évènement\n")
	fmt.Println("5. Rechercher un évènement\n")
	fmt.Println("6. Quitter\n")
	fmt.Println("Chosissez une option : ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		/* TODO: Créer un nouvel évènement */
		break

	case 2:
		/* TODO: Visualiser les évènements */
		break

	case 3:
		/* TODO: Modifier un évènement */
		break

	case 4:
		/* TODO: Supprimer un évènement */
		break

	case 5:
		/* TODO: Recherchez un évènment */
		break

	case 6:
		/* TODO: Quitter */
		break

	default:
		fmt.Println("Choix invalide")
		menu()
		break
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

	//Gestion d'erreur en cas d'heure antérieure à l'heure actuelle
	if parsedHour.Before(time.Now()) {
		fmt.Println("Heure invalide")
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

}
