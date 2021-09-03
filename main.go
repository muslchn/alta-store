package main

import (
	"alta-store/config"
	routes "alta-store/route"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
