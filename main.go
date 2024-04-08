package main

import (
	"flag"
	"fmt"
)

const VERSION = "1.0"

func main() {
	version := flag.Bool("version", false, "Affiche la version du programme")
	flag.Parse()

	if *version {
		fmt.Println(VERSION)
		return
	}
}
