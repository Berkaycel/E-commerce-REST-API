package main

import (
	"ecommerce-api/config"
	"ecommerce-api/models"
	"ecommerce-api/routes"
	"fmt"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{})

	fmt.Println("Database migrated successfully!")

	r := routes.SetupRoutes()

	fmt.Println("Server is running on port 8080...")
	r.Run(":8080")
}
