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
	fmt.Println("Servidor corriente en el puerto", port)
	r.Run(":" + port)
}
