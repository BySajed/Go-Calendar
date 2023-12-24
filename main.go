package main

import "fmt"

func menu() {
	choice := 0

	fmt.Println("Bienvenue dans le système de gestion de planning")
	fmt.Println("1. Créer un nouvel évènement")
	fmt.Println("2. Visualiser les évènements")
	fmt.Println("3. Modifier un évènement")
	fmt.Println("4. Supprimer un évènement")
	fmt.Println("5. Rechercher un évènement")
	fmt.Println("6. Quitter")
	fmt.Println("Chosissez une option : ")

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
