package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func scanString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func newEvent() {
	var title string
	var date string
	fmt_date := "2006-01-02"
	var hour string
	var place string
	var category string
	var description string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("Entrez le titre de l'évènement : ")
	title = scanString(scanner)
	fmt.Println("Entrez la date de l'évènement (YYYY-MM-DD): ")
	date = scanString(scanner)

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
	hour = scanString(scanner)

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
	place = scanString(scanner)

	fmt.Println("Choisissez une catégorie(Professionnel, Personnel, Loisir): ")
	for {
		category = scanString(scanner)
		if category != "Professionnel" && category != "Personnel" && category != "Loisir" {
			fmt.Println("Catégorie invalide")
			continue
		}
		break
	}

	fmt.Println("Entrez une brève Description de l'évènement : ")
	description = scanString(scanner)

	newEvent := Event{
		ID:          idEvent,
		Title:       title,
		Date:        date,
		Hour:        hour,
		Place:       place,
		Category:    category,
		Description: description,
	}
	if db.Ping() == nil {
		newEvent = NewEventRepository(newEvent)
	}

	eventsMap[idEvent] = newEvent
	idEvent++
	fmt.Println("Evènement créé avec succès !")
}

func NewEventRepository(event Event) Event {
	query := `INSERT INTO event (title, date, hour, place, category, description) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	session, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return event
	}
	defer func(session *sql.Stmt) {
		err := session.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(session)

	err = session.QueryRow(event.Title, event.Date, event.Hour, event.Place, event.Category, event.Description).
		Scan(&event.ID)
	if err != nil {
		log.Fatal(err)
		return event
	} else {
		log.Println("événement ajouté avec succès !!!")
	}

	return event
}
