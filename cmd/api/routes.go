package main

import (
	"database/sql"
	"net/http"

	"github.com/codepnw/go-movie-booking/internal/handlers"
	"github.com/codepnw/go-movie-booking/internal/repositories"
	"github.com/codepnw/go-movie-booking/internal/services"
	"github.com/gin-gonic/gin"
)

func apiRoutes(r *gin.Engine, db *sql.DB, version string) {
	router := r.Group("/" + version)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK!"})
	})

	userRepo := repositories.NewUserRepository(db)
	userSrv := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userSrv)
	// User routes
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
	router.GET("/profile/:id", userHandler.GetProfile)
}
