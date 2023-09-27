package main

import (
	_ "blogpost/db_connection"
	"blogpost/route"
	"os"

	"github.com/joho/godotenv"
	logging "github.com/sirupsen/logrus"
)

func main() {
	route.InitializeRouter()
	err := godotenv.Load(".env")
	if err != nil {
		logging.Error("Error while loading .env file")
		os.Exit(1)
	}
}
