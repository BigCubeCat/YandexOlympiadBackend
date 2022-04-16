package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"search/api"
	"search/db"
)

func setupLog() {
	logFileName := os.Getenv("LOG_FILE")
	path := logFileName
	_ = os.Remove(path)
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Unable to open log file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panicf("Unable to close log file: %v", err)
		}
	}(f)
	log.SetOutput(f)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	setupLog()
	db.InitDB(os.Getenv("DATABASE"))
	address := os.Getenv("ADDRESS")
	host := os.Getenv("HOST")

	app := fiber.New(fiber.Config{})

	app.Add("GET", "/find", api.GetRecipes)
	app.Add("POST", "/add", api.CreateRecipe)

	log.Fatal(app.Listen(host + ":" + address))

}
