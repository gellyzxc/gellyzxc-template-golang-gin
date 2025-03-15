package main

import (
	"fmt"
	"log"
	"os"

	"gellyzxc-template-golang-gin/config"
	"gellyzxc-template-golang-gin/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)

	fmt.Println("port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("error:", err)
	}
}
