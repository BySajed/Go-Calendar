package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Event struct {
	ID          int
	Title       string
	Date        string
	Hour        string
	Place       string
	Category    string
	Description string
}

var eventsMap = make(map[int]Event)
var idEvent = 1

func menu() {
	choice := 0
	cExit := 0

	fmt.Println("\n------------------------------------------------\n")
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
			menu()

		case 2:
			showEvents()
			choice = 0
			menu()

		case 3:
			modifyEvent()
			choice = 0
			menu()

		case 4:
			/* TODO: Supprimer un évènement */
			choice = 0
			menu()

		case 5:
			/* TODO: Recherchez un évènment */
			choice = 0
			menu()

		case 6:
			/* TODO: Quitter */
			choice = 0
			cExit++
		default:
			fmt.Println("Entrée invalide, veuillez réessayez")
		}
	}
}

func loadEnvFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
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
	// Fermeture de la base de donnée
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Erreur fermeture de la base de données")
		}
	}(db)

	fmt.Println("Bienvenue dans le système de gestion de planning\n")
	connectDatabase()
	menu()
}
