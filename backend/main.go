package main

import (
	"LostPaws/backend/config"
	"LostPaws/backend/routes"
	"fmt"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	port := "8080"
	fmt.Printf("Servidor corriente en el puero", port)
	r.Run(":", port)
}
