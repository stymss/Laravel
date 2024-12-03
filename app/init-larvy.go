package main

import (
	"log"
	"os"

	larvy "github.com/stymss/Laravel"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	lrvl := &larvy.Larvy{}
	err = lrvl.New(path)
	if err != nil {
		log.Fatal(err)
	}

	lrvl.AppName = "Application"
	lrvl.Debug = false
	lrvl.Version = "1.0.0"

	app := &application{
		App: lrvl,
	}

	return app

}
