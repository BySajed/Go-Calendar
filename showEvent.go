package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	EVERYTHING = iota
	FILTER
	SORT
)

func copyEvents() []Event {
	var dest []Event

	for _, v := range eventsMap {
		dest = append(dest, v)
	}
	return dest
}

func filterDate() {

}

func filterMain() {

}

func choiceMode() int {
	var choiceStr string
	choice := -1
	var err error
	var err2 error

	for choice < 0 || choice > 2 || err != nil || err2 != nil {
		fmt.Print("Veuillez choisir votre mode\n1 - Tous les événéments\n2 - Filtrer la liste\n3 - Trier la liste\n")
		_, err = fmt.Scan(&choiceStr)
		choice, err2 = strconv.Atoi(choiceStr)
		choice--
	}
	return choice
}

func fetchEventsRepository() {
	rows, err := db.Query("SELECT * from event")
	if err != nil {
		log.Fatal("Erreur de récupération des données")
		return
	}

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Title, &event.Date, &event.Hour, &event.Place, &event.Category, &event.Description)
		if err != nil {
			log.Fatal("Fatal erreur de récupération des données")
			return
		}
		eventsMap[event.ID] = event
	}
}

func showEvents() {

	if db.Ping() == nil {
		fetchEventsRepository()
	}

	fmt.Println("Voulez-vous voir tous vos évènements ou un en particulier ? (Tous/Un)")
	var choice string
	fmt.Scan(&choice)

	if choice == "Tous" {
		mode := choiceMode()
		var eventShow []Event
		switch mode {
		case EVERYTHING:
			eventShow = copyEvents()
			//case FILTER:

		}
		fmt.Println("Voici la liste des évènements : \n")
		for i := 0; i < len(eventShow); i++ {
			fmt.Println("------------------------------------------------\n")
			fmt.Println("Evènement n°", i+1)
			fmt.Println("Titre : ", eventShow[i].Title)
			fmt.Println("Date : ", eventShow[i].Date)
			fmt.Println("Heure : ", eventShow[i].Hour)
			fmt.Println("Lieu : ", eventShow[i].Place)
			fmt.Println("Catégorie : ", eventShow[i].Category)
			fmt.Println("Description : ", eventShow[i].Description)
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
