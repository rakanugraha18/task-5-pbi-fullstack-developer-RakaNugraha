package main

import (
	"fmt"
	"log"
	"task_5_pbi_btpns_RakaNugraha/database"
	"task_5_pbi_btpns_RakaNugraha/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Connect to the database
	_, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Informasi sukses koneksi
	fmt.Println("Successfully connected to database")

	// Setup router
	r := router.SetupRouter()

	// Jalankan server
	r.Run(":8080")
}
