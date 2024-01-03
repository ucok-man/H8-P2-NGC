package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/app"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/model"
)

func init() {
	godotenv.Load()
}

func main() {
	app := app.New()

	// setup migration
	if err := app.Config.GetDBConn().AutoMigrate(&model.Store{}, &model.Product{}); err != nil {
		log.Fatal("Error Auto Migrate: ", err)
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
