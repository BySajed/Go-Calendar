package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
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

func choiceCategorie() string {
	choice := ""
	for choice != "1" && choice != "2" && choice != "3" {
		fmt.Println("Quelle catégorie ?\n1 - Professionnel\n2 - Personnel\n3 - Loisir\n4 - Non, j'ai fait une erreur")
		fmt.Scan(&choice)
		if choice == "4" {
			return "4"
		}
	}
	switch choice {
	case "1":
		choice = "Professionnel"
	case "2":
		choice = "Personnel"
	case "3":
		choice = "Loisir"
	}
	return choice
}

func filterCategorie(events []Event) []Event {
	choice := choiceCategorie()
	if choice == "4" {
		return events
	}

	var newEvent []Event
	for i := 0; i < len(events); i++ {
		if events[i].Category == choice {
			newEvent = append(newEvent, events[i])
		}
	}
	return newEvent
}

func filterMain(events []Event) []Event {
	choice := ""

	for choice != "Y" && choice != "N" {
		fmt.Println("Voulez-vous filter par catégorie ? (Y/N)")
		fmt.Scan(&choice)
		if choice == "Y" {
			events = filterCategorie(events)
		}
	}
	return events
}

func compareDate(event1 Event, event2 Event) bool {
	date1, _ := time.Parse("2006-01-02T15:04:05Z", event1.Date)
	date2, _ := time.Parse("2006-01-02T15:04:05Z", event2.Date)
	return date1.Before(date2)
}

func sortDate(events []Event) []Event {
	sort.Slice(events, func(i, j int) bool {
		return compareDate(events[i], events[j])
	})
	return events
}

func sortCategorie(events []Event) []Event {
	choice := choiceCategorie()
	if choice == "4" {
		return events
	}

	var newEvents []Event
	for i := 0; i < len(events); i++ {
		if events[i].Category == choice {
			newEvents = append(newEvents, events[i])
		}
	}
	for i := 0; i < len(events); i++ {
		if events[i].Category != choice {
			newEvents = append(newEvents, events[i])
		}
	}

	return newEvents
}

func sortMain(events []Event) []Event {
	choice := ""
	for choice != "1" && choice != "2" && choice != "3" {
		fmt.Println("Voulez-vous trier par date ou par catégorie ?\n1 - date\n2 - catégorie\n3 - Non, j'ai fait une erreur")
		fmt.Scan(&choice)
	}

	if choice == "1" {
		events = sortDate(events)
	} else if choice == "2" {
		events = sortCategorie(events)
	}

	return events
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

func showAEvent(event Event) {
	fmt.Println("------------------------------------------------")
	fmt.Println("Evènement n°", event.ID)
	fmt.Println("Titre : ", event.Title)
	fmt.Println("Date : ", event.Date)
	fmt.Println("Heure : ", event.Hour)
	fmt.Println("Lieu : ", event.Place)
	fmt.Println("Catégorie : ", event.Category)
	fmt.Println("Description : ", event.Description)
}

func showEvents() {

	if db.Ping() == nil {
		fetchEventsRepository()
	}

	for {
		fmt.Println("Voulez-vous voir tous vos évènements ou un en particulier ? (Tous/Un)")
		var choice string
		fmt.Scan(&choice)

		if choice == "Tous" {
			mode := choiceMode()
			eventShow := copyEvents()
			switch mode {
			case FILTER:
				eventShow = filterMain(eventShow)
			case SORT:
				eventShow = sortMain(eventShow)

			}
			fmt.Println("Voici la liste des évènements : \n")
			for i := 0; i < len(eventShow); i++ {
				fmt.Println("------------------------------------------------")
				fmt.Println("Evènement n°", eventShow[i].ID)
				fmt.Println("Titre : ", eventShow[i].Title)
				fmt.Println("Date : ", eventShow[i].Date)
				fmt.Println("Heure : ", eventShow[i].Hour)
				fmt.Println("Lieu : ", eventShow[i].Place)
				fmt.Println("Catégorie : ", eventShow[i].Category)
				fmt.Println("Description : ", eventShow[i].Description)
			}
		} else if choice == "Un" {
			var id int
			for {
				fmt.Println("Entrez le numéro de l'évènement : ")
				fmt.Scan(&id)
				if _, i := eventsMap[id]; !i {
					fmt.Println("L'événement n'existe pas")
					continue
				}
				break
			}
			showAEvent(eventsMap[id])
		} else {
			fmt.Println("Choix invalide")
			continue
		}
		break
	}
	return
}
