package main

import (
	"encoding/csv"
	_ "encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func askName() string {
	var nameFile string
	fmt.Println("Quel nom doit avoir le fichier ?")
	for {
		_, err := fmt.Scanln(&nameFile)
		if err != nil {
			fmt.Println("Entrée invalide")
			continue
		}
		if len(nameFile) >= 249 && len(nameFile) < 1 {
			fmt.Println("Nom de fichier invalide")
			continue
		}
		break
	}
	return nameFile
}

func exportJSON() {
	if db.Ping() == nil {
		fetchEventsRepository()
	}

	nameFile := askName()

	json, err := json.MarshalIndent(eventsMap, "", " ")
	if err != nil {
		log.Fatal("Erreur lors de la conversion en JSON:", err)
		return
	}

	file, err := os.Create(nameFile + ".json")
	if err != nil {
		log.Fatal("Erreur d'ouverture du fichier :", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Erreur de fermeture du fichier :", err)
			return
		}
	}(file)

	_, err = file.Write(json)
	if err != nil {
		log.Fatal("Erreur lors de l'écriture dans le fichier:", err)
		return
	}

	fmt.Println("Les données ont été exporté en JSON !!!")
}

func exportCSV() {
	if db.Ping() == nil {
		fetchEventsRepository()
	}

	nameFile := askName()
	file, err := os.Create(nameFile + ".csv")
	if err != nil {
		log.Fatal("Erreur d'ouverture du fichier :", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Erreur de fermeture du fichier :", err)
			return
		}
	}(file)
	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := []string{"id", "title", "date", "hour", "place", "description"}
	writer.Write(header)
	for key, event := range eventsMap {
		colonne := []string{
			strconv.Itoa(key),
			event.Title,
			event.Date,
			event.Hour,
			event.Place,
			event.Category,
			event.Description,
		}
		writer.Write(colonne)
	}

	fmt.Println("Les données ont été exporté en CSV !!!")
}
