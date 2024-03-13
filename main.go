package main

import (
	"assignment-two/database"
	"assignment-two/routers"
)

func main() {
	//Koneksi database
	database.StartDB()

	//Setup Router (gin engine)
	rts := routers.SetupRouter()

	rts.Run()
}
