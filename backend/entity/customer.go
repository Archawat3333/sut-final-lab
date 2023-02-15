package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `gorm:"primaryKed" valid:"required~Name cannot be null"`
	Email      string 
	CustomerID string   `gorm:"primaryKed" valid:"matches(^([L|M|H])([0-9]{7}))~Customerid not correct format"`
}
