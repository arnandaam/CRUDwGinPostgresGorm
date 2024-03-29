package main

import (
	"showcase/app"
	"showcase/config"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = config.InitPostgres()
	if err != nil {
		panic(err)
	}

	err = config.InitGorm()
	if err != nil {
		panic(err)
	}

}

func main() {
	app.StartApplication()
}
