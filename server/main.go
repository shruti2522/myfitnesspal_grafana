package main

import (
	"os"
	"fmt"

	"myfitnesspal-grafana/routers"
	"myfitnesspal-grafana/database"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	router := gin.Default()

	router.Use(cors.Default())

	db, err := database.InitDB()

	if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
    }
    defer db.Close()


	routers.InitRoutes(router)

	router.Run(":" + port)
}
