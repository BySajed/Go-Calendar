package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Event struct {
	ID          int    `csv:"id"`
	Title       string `csv:"title"`
	Date        string `csv:"date"`
	Hour        string `csv:"hour"`
	Place       string `csv:"place"`
	Category    string `csv:"category"`
	Description string `csv:"description"`
}

var eventsMap = make(map[int]Event)
var idEvent = 1
var db *sql.DB

func notification() {
	if db.Ping() == nil {
		fetchEventsRepository()
	}
	numberEventToday := 0
	for key, event := range eventsMap {
		eventTime, err := time.Parse("2006-01-02T15:04:05Z", event.Date)
		if err != nil {
			log.Fatal(err)
		}
		today := time.Now().UTC().Truncate(24 * time.Hour)

		eventTime = eventTime.Truncate(24 * time.Hour)
		if eventTime.Equal(today) {
			if numberEventToday == 0 {
				fmt.Println("Liste des événements se déroulant aujourd'hui :")
			}
			fmt.Printf("N°%d : %s à %s\n", key, event.Title, event.Hour)
			numberEventToday++
		}
	}
	if numberEventToday == 0 {
		fmt.Println("Aucun événément se déroulant aujourd'hui")
	} else {
		fmt.Printf("Il y a %d événements se déroulant aujourd'hui", numberEventToday)
	}
}

func menu() {
	choice := 0
	cExit := 0

	notification()
	for cExit != 1 {
		var line string
		errorInput := true
		fmt.Println("\n------------------------------------------------")
		fmt.Println("1. Créer un nouvel évènement")
		fmt.Println("2. Visualiser les évènements")
		fmt.Println("3. Modifier un évènement")
		fmt.Println("4. Supprimer un évènement")
		fmt.Println("5. Exporter la base en JSON")
		fmt.Println("6. Exporter la base en CSV")
		fmt.Println("7. Quitter")
		fmt.Println("Chosissez une option : ")

		for errorInput {
			errorInput = false
			fmt.Scanln(&line)
			choice, _ = strconv.Atoi(line)

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
				deleteEvent()
				choice = 0
			case 5:
				exportJSON()
				choice = 0
			case 6:
				exportCSV()
				choice = 0
			case 7:
				/* TODO: Quitter */
				choice = 0
				cExit++
			default:
				fmt.Println("Entrée invalide, veuillez réessayez")
				errorInput = true
			}
		}
	}
}

func loadEnvFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Fichier 'config.env' introuvable")
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Erreur fermeture fichier env")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		if len(parts) == 2 {
			err = os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func connectDatabase() {
	err := loadEnvFromFile("config.env")

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName)

	db, err = sql.Open("postgres", connect)
	if err != nil {
		log.Println("échec de la connexion à la base de données")
		log.Fatal(err)
	} else {
		log.Println("Connexion à la base de données réussie ")
	}

}

func main() {
	fmt.Println("Bienvenue dans le système de gestion de planning\n")
	connectDatabase()
	defer db.Close()
	menu()
}
