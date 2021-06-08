package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty"`
	Name             string `json:"name,omitempty"`
	Email            string `json:"email,omitempty"`
	Cakes            Cake   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
