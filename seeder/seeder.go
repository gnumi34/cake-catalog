package main

import (
	"cake-catalog/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("user=postgres password=12345678 host=localhost port=5432 dbname=catalog sslmode=disable"))
	if err != nil {
		panic(err.Error())
	}

	password := "iamspeed"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Something is wrong when hashing the password!")
	}

	user1 := models.User{
		Username: "gnumi34",
		Password: string(hash),
		Name:     "Tripamungkas Kartika Adi",
		Email:    "tripamungkas.adi@gmail.com",
	}

	existingUser := models.User{}
	err = db.First(&existingUser, 1).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic("Something is wrong when fetching user!")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.Create(&user1).Error
	}

	if err != nil {
		panic("something is wrong when creating user!")
	}
}
