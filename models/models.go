package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string
	Password         string
	Name             string
	Email            string
	Cakes            Cake `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TokenReset       string
	TokenResetExpiry time.Time
}

type Cake struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}
