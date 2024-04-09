package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Erreur survenue à %v : %s", e.When, e.What)
}

func run() error {
	return MyError{
		When: time.Now(),
		What: "Une erreur s'est produite!",
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
