package main

import (
	"alta-store/config"
	"alta-store/lib/database"
	"alta-store/routes"
)

func main() {
	config.InitDB()
	database.InitData()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
