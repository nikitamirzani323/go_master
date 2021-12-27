package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nikitamirzani323/go_mastertoto/routers"
)

func main() {
	var errenv = godotenv.Load()
	if errenv != nil {
		panic("Failed to load env file")
	}
	app := routers.Init()
	log.Fatal(app.Listen(":7072"))
}
