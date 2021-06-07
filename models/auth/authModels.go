package auth

import (
	"cake-catalog/database"
	"cake-catalog/models"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LogIn(authData models.AuthRequest) (models.User, error) {
	user := models.User{}
	err := database.Conn.Where("username = ?", authData.Username).First(&user).Error

	log.Println("[INFO] User:", user)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user not found")
		} else {
			return user, err
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authData.Password)); err != nil {
		return user, fmt.Errorf("incorrect password")
	}

	return user, nil
}