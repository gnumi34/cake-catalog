package auth

import (
	"cake-catalog/database"
	"cake-catalog/models"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LogIn(authData models.AuthRequest) (models.User, error) {
	user := models.User{}
	err := database.Conn.Where("username = ?", authData.Username).First(&user).Error

	log.Println("[INFO] User:", user)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user is not found")
		} else {
			return user, fmt.Errorf("error on LogIn: %s", err.Error())
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authData.Password)); err != nil {
		return user, fmt.Errorf("incorrect password")
	}

	return user, nil
}

func SignUp(userData models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error on bcrypt: %s", err.Error())
	}

	userData.Password = string(hashedPassword)

	err = database.Conn.Create(&userData).Error
	if err != nil {
		return fmt.Errorf("error on DB creation: %s", err.Error())
	}

	return nil
}

func ForgetPassword(username string) (string, error) {
	getUser := models.User{}
	err := database.Conn.Where("username = ?", username).First(&getUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("user is not found")
		} else {
			return "", fmt.Errorf("error on ForgetPassword: %s", err.Error())
		}
	}

	token := sha256.New()
	_, err = token.Write([]byte(getUser.Password + getUser.CreatedAt.String()))
	if err != nil {
		return "", fmt.Errorf("error on ForgetPassword: %s", err.Error())
	}

	getUser.TokenReset = base64.URLEncoding.EncodeToString(token.Sum(nil))
	getUser.TokenResetExpiry = time.Now().Add(time.Duration(15) * time.Minute)

	err = database.Conn.Save(&getUser).Error
	if err != nil {
		return "", fmt.Errorf("error on ForgetPassword: %s", err.Error())
	}

	return getUser.TokenReset, nil
}
