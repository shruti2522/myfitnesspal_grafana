package models

import(
	"github.com/jinzhu/gorm"
)

type UserEntry struct {
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