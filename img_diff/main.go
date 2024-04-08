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
		resp, err := http.Get(url)
		if err != nil {
			logger.Printf("Erreur lors du téléchargement de %s : %v\n", url, err)
			continue
		}
		defer resp.Body.Close()

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

		images = append(images, filename)
	}

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

	if uniqueImage != "" {
		fmt.Printf("L'image unique est : %s\n", uniqueImage)
	} else {
		logger.Println("Aucune image unique trouvée.")
	}
}
