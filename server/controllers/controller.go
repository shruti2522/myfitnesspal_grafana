package controllers

import (
	"myfitnesspal-grafana/database"
	"myfitnesspal-grafana/utils"

	"github.com/gin-gonic/gin"
)

func UploadCSV(context *gin.Context) {
	
	db, err := database.InitDB()

    file, err := context.FormFile("file")
    if err != nil {
        context.String(400, "Bad Request")
        return
    }

    if err := context.SaveUploadedFile(file, "uploads/"+file.Filename); err != nil {
        context.String(500, "Internal Server Error")
        return
    }

    if err := utils.ProcessCSV(db,"uploads/" + file.Filename); err != nil {
        context.String(500, "Error processing csv file")
        return
    }

    context.String(200, "File uploaded successfully")
}