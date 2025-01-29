package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/neerajbg/go-gin-auth/database"
	"github.com/neerajbg/go-gin-auth/routes"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error in loading env file.")
	}
	database.ConnectDB()
}

func main() {
	sqlDb, err := database.DBConn.DB()
	if err != nil {
		log.Println("Error in getting db conn.")
	}
	defer sqlDb.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Static file serving
	router.Static("/static", "./static") // Thêm dòng này để phục vụ file tĩnh

	// Middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders: []string{"Origin", "Auth-token", "token", "Content-type"},
	}))

	routes.SetupRoutes(router)

	log.Fatal(router.Run(":" + port))
}
