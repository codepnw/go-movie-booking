package main

import (
	"log"
	"os"

	"github.com/codepnw/go-movie-booking/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed loading .env: %v", err)
	}

	db, err := database.InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed connection the database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	apiRoutes(router, db, "v1")

	router.Run(":8080")
}
