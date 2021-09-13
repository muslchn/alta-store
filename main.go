package main

import (
	"alta-store/config"
	"alta-store/lib/database"
	routes "alta-store/route"
)

func main() {
	config.InitDB()
	database.InitData()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
