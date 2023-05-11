package main

import (
	"Mini_Project_Toko-Online/configs"
	"Mini_Project_Toko-Online/routes"
)

func main() {
	// create a new echo instance
	configs.InitDB()
	configs.InitialMigration()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
