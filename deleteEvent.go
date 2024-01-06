package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func deleteEvent() {
	if db.Ping() == nil {
		fetchEventsRepository()
	}

	var idChoice int
	var line string
	var err error
	fmt.Println("Quel est l'id de l'événement que vous voulez supprimer ?")
	for {
		fmt.Scanln(&line)
		idChoice, err = strconv.Atoi(line)
		if err != nil {
			fmt.Println("L'entrée est invalide")
			continue
		}
		if _, i := eventsMap[idChoice]; !i {
			fmt.Println("L'id n'existe pas")
			continue
		}
		break
	}
	showAEvent(eventsMap[idChoice])
	fmt.Println("-------------------")
	fmt.Println("êtes-vous sûr de vouloir supprimer cette événement ? (Y/N)")
	for {
		fmt.Scanln(&line)
		if line == "Y" {
			break
		} else if line == "N" {
			return
		} else {
			fmt.Println("Choix invalide")
		}
	}
	delete(eventsMap, idChoice)
	deleteEventRepository(idChoice)
}

func deleteEventRepository(id int) {
	query := `DELETE FROM event WHERE id=$1`
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

	_, err = session.Exec(id)
	if err != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("événement supprimé avec succès !!!")
	}
}
