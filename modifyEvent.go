package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func modifyEvent() {
	if db.Ping() == nil {
		fetchEventsRepository()
	}

	var line string
	var id int
	var err error

	fmt.Println("Entrez le numéro de l'évènement que vous souhaitez modifier : ")
	for {
		fmt.Scanln(&line)
		id, err = strconv.Atoi(line)
		if err != nil {
			fmt.Println("L'entrée est invalide")
			continue
		}
		if _, i := eventsMap[id]; !i {
			fmt.Println("L'id n'existe pas")
			continue
		}
		break
	}
	event := eventsMap[id]
	showAEvent(event)

	for {
		fmt.Println("Que souhaitez-vous modifier ? (Titre/Date/Heure/Lieu/Catégorie/Description)")
		_, err := fmt.Scanln(&line)
		if err != nil {
			continue
		}

		if line == "Titre" || line == "titre" {
			fmt.Println("Entrez le nouveau titre : ")
			var newTitle string
			fmt.Scanln(&newTitle)
			event.Title = newTitle
			fmt.Println("Titre modifié avec succès !")

		} else if line == "Date" || line == "date" {
			fmt.Println("Entrez la nouvelle date : ")
			var newDate string
			for {
				fmt.Println("Entrez la date de l'évènement (YYYY-MM-DD): ")
				fmt.Scanln(&newDate)
				newDate = strings.TrimSpace(newDate)

				//Gestion d'erreur en cas de mauvais format de date
				parsedTime, err := time.Parse("2006-01-02", newDate)
				if err != nil {
					fmt.Println("Date invalide")
					continue
				}
				if parsedTime.Format("2006-01-02") != newDate {
					fmt.Println("e format date est invalide")
					continue
				}
				break
			}
			event.Date = newDate
			fmt.Println("Date modifiée avec succès !")

		} else if line == "Heure" || line == "heure" {
			fmt.Println("Entrez la nouvelle heure : ")
			var newHour string
			for {
				fmt.Println("Entrez l'heure de l'évènement (HH:MM): ")
				fmt.Scanln(&newHour)
				newHour = strings.TrimSpace(newHour)

				//Gestion d'erreur en cas de mauvais format d'heure
				parsedHour, err := time.Parse("15:04", newHour)
				if err != nil {
					fmt.Println("Heure invalide")
					continue
				}
				if parsedHour.Format("15:04") != newHour {
					fmt.Println("Le format heure est invalide")
					continue
				}
				break
			}
			event.Hour = newHour
			fmt.Println("Heure modifiée avec succès !")

		} else if line == "Lieu" || line == "lieu" {
			fmt.Println("Entrez le nouveau lieu : ")
			var newPlace string
			fmt.Scanln(&newPlace)
			event.Place = newPlace
			fmt.Println("Lieu modifié avec succès !")

		} else if line == "Catégorie" || line == "catégorie" {
			fmt.Println("Entrez la nouvelle catégorie : ")
			var newCategory string
			fmt.Scanln(&newCategory)
			event.Category = newCategory
			fmt.Println("Catégorie modifiée avec succès !")

		} else if line == "Description" {
			fmt.Println("Entrez la nouvelle Description : ")
			var newDescription string
			fmt.Scanln(&newDescription)
			event.Description = newDescription
			fmt.Println("Description modifiée avec succès !")
		} else {
			fmt.Println("Choix invalide")
			continue
		}

		for line != "Y" && line != "N" {
			fmt.Println("Voulez-vous modifier autre chose dans l'événement ? (Y/N)")
			fmt.Scanln(&line)
		}
		if line == "Y" {
			showAEvent(event)
			continue
		}
		break
	}

	if db.Ping() == nil {
		modifyEventRepository(event)
	}
	eventsMap[id] = event
}

func modifyEventRepository(event Event) {
	query := `UPDATE event SET title = $1, date = $2, hour = $3, place = $4, category = $5, description = $6 WHERE id = $7`
	session, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(session *sql.Stmt) {
		err := session.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(session)

	_, err = session.Exec(event.Title, event.Date, event.Hour, event.Place, event.Category, event.Description, event.ID)
	if err != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("événement modifié avec succès !!!")
	}
}
