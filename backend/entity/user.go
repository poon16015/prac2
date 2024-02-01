package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `valid:"required~Name is required"`
	Email string `valid:"required~Email is required,email~Email is invalid"`
	Phone string `valid:"required~phone is required, stringlength(10|10)~PhoneNumber length is not 10 digits."`
}