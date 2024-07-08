package models

import (
	"gorm.io/gorm"
	"time"
)

type Statement struct {
	gorm.Model
	Account     string    `gorm:"column:Account"`
	Date        time.Time `gorm:"column:Date"`
	Income      float32   `gorm:"column:Income"`
	Expenditure float32   `gorm:"column:Expenditure"`
}
