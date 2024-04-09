package main

import "fmt"

func PrintIt(input interface{}) {
	switch v := input.(type) {
	case int:
		fmt.Printf("Le type de l'argument est int, valeur: %v\n", v)
	case string:
		fmt.Printf("Le type de l'argument est string, valeur: %v\n", v)
	default:
		fmt.Printf("Le type de l'argument est %T, valeur: %v\n", v, v)
	}
}

func main() {
	PrintIt(42)
	PrintIt("Hello, World!")
	PrintIt(true)
}
