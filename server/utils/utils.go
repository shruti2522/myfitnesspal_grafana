package utils

import(
	"encoding/csv"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "myfitnesspal-grafana/models"
	"github.com/jinzhu/gorm"
)

func ProcessCSV(db *gorm.DB,filePath string) error {
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

		entry := models.UserEntry{
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