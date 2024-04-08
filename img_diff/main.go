package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "", log.LstdFlags)

	images := []string{"image_1.jpg", "image_2.jpg", "image_3.jpg"}
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
