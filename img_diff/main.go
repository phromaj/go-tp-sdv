package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Création d'un logger pour écrire les messages d'erreur dans stderr
	logger := log.New(os.Stderr, "", log.LstdFlags)

	// Ajout du flag pour passer une liste d'URLs en argument
	urlsFlag := flag.String("urls", "", "Liste d'URLs séparées par des virgules")
	flag.Parse()

	// Parsing de la chaîne d'URLs
	var urls []string
	if *urlsFlag != "" {
		urls = strings.Split(*urlsFlag, ",")
	}

	// Téléchargement des images depuis les URLs
	images := []string{}
	for i, url := range urls {
		// Téléchargement de l'image depuis l'URL
		resp, err := http.Get(url)
		if err != nil {
			logger.Printf("Erreur lors du téléchargement de %s : %v\n", url, err)
			continue
		}
		defer resp.Body.Close()

		// Lecture des données de l'image
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Printf("Erreur lors de la lecture de %s : %v\n", url, err)
			continue
		}

		// Écriture de l'image sur le disque
		filename := fmt.Sprintf("image_%d.jpg", i+1)
		err = os.WriteFile(filename, data, 0644)
		if err != nil {
			logger.Printf("Erreur lors de l'écriture de %s : %v\n", filename, err)
			continue
		}

		// Ajout du nom de fichier à la liste des images téléchargées
		images = append(images, filename)
	}

	// Calcul des hashes SHA-256 pour chaque image
	hashes := make([]string, len(images))
	for i := 0; i < len(images); i++ {
		file, err := os.Open(images[i])
		if err != nil {
			logger.Printf("Erreur lors de l'ouverture de %s : %v\n", images[i], err)
			return
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		hash := sha256.New()
		_, err = reader.WriteTo(hash)
		if err != nil {
			logger.Printf("Erreur lors de la lecture de %s : %v\n", images[i], err)
			return
		}

		hashes[i] = fmt.Sprintf("%x", hash.Sum(nil))
	}

	// Recherche de l'image unique
	var uniqueImage string
	for i := 0; i < len(images); i++ {
		isDuplicate := false
		for j := 0; j < len(images); j++ {
			if i != j && hashes[i] == hashes[j] {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			uniqueImage = images[i]
			break
		}
	}

	// Affichage de l'image unique ou d'un message d'erreur si aucune image unique n'est trouvée
	if uniqueImage != "" {
		fmt.Printf("L'image unique est : %s\n", uniqueImage)
	} else {
		logger.Println("Aucune image unique trouvée.")
	}
}
