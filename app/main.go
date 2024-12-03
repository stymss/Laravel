package main

import (
	"fmt"

	larvy "github.com/stymss/Laravel"
)

type application struct {
	App *larvy.Larvy
}

func main() {
	fmt.Println("Welcome to application")
	initApplication()
}
