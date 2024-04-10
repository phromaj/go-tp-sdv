package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type User struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	UserID   string `json:"userID"`
}

var userMap map[string]User

func main() {
	// Chargement des utilisateurs depuis le fichier JSON
	loadUsersFromFile()

	// Création du serveur HTTP
	http.HandleFunc("/", userHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadUsersFromFile() {
	// Lecture du fichier JSON
	data, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal("Error reading users file:", err)
	}

	// Désérialisation du JSON dans un slice d'utilisateurs
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	// Remplissage de la map avec les utilisateurs
	userMap = make(map[string]User)
	for _, user := range users {
		userMap[user.UserID] = user
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération du userID depuis la query string
	userID := r.FormValue("id")

	// Recherche de l'utilisateur correspondant dans la map
	user, found := userMap[userID]

	// Définition du content-type de la réponse
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if !found {
		// Utilisateur non trouvé
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Utilisateur trouvé
	w.WriteHeader(http.StatusOK)

	// Sérialisation de l'utilisateur en JSON et envoi de la réponse
	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Error serializing user to JSON:", err)
	}
	w.Write(jsonUser)
}
