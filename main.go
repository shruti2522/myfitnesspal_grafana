package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=myfitnesspal_csv password=admin sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}
	db.AutoMigrate(&userEntry{})
}

type userEntry struct {
	gorm.Model
	Date         string
	Entry        string
	Item         string
	Amount       string
	Calories     float64
	Carbs        string
	Fat          string
	Protein      string
	Cholest      string
	Sodium       string
    Sugars       string
	Fiber        string
}

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(400, "Bad Request")
			return
		}

		if err := c.SaveUploadedFile(file, "uploads/"+file.Filename); err != nil {
			c.String(500, "Internal Server Error")
			return
		}

		if err := processCSV("uploads/" + file.Filename); err != nil {
			c.String(500, "Error uploading the csv file")
			return
		}

		c.String(200, "File uploaded successfully")
	})

	r.Run(":8080")
}

func processCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		entry := userEntry{
			Date:         record[0],
	        Entry:        record[1],
	        Item:         record[2],
			Amount:       record[3],
			Calories:     parseFloat(record[4]),
			Carbs:        record[5],
			Fat:          record[6],
			Protein:      record[7],
			Cholest:      record[8],
			Sodium:       record[9],
			Sugars:       record[10],
			Fiber:        record[11],
		}

		if err := db.Create(&entry).Error; err != nil {
			fmt.Println("Error inserting entry into the database:", err)
		}
	}

	return nil
}

func parseFloat(s string) float64 {
	val, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0.0
	}
	return val
}