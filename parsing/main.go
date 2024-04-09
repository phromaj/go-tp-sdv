package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Login    string `json:"userName"`
	Password string
}

func main() {
	// Création d'un utilisateur
	user := User{
		Login:    "Paul",
		Password: "pass123",
	}

	// Sérialisation de l'utilisateur en JSON
	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation :", err)
		return
	}

	// Affichage du JSON de l'utilisateur dans le terminal
	fmt.Println(string(jsonUser))

	// Lecture du fichier users.json
	data, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Désérialisation du contenu du fichier dans une liste de User
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation :", err)
		return
	}

	// Affichage de la liste des utilisateurs
	fmt.Println(users)
}
