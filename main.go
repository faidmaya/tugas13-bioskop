package main

import (
	"tugas13-bioskop/database"
	"tugas13-bioskop/routers"
)

func main() {
	database.ConnectDB()

	r := routers.StartServer()

	r.Run(":8080")
}
