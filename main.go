package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"verralive/controllers"
	"verralive/models"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	models.ConnectDatabase()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{os.Getenv("APP_URL")}
	corsConfig.AddAllowMethods("OPTIONS")

	r := gin.Default()

	r.Use(cors.New(corsConfig))

	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Presentation Management Systems API Online."})
	})

	public.POST("/email", controllers.CreateEmail)

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		fmt.Println("Error starting server.")
		return
	}
}
