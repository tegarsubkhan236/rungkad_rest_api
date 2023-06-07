package service

import (
	"errors"
	"example/pkg/config"
	"example/pkg/model"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)

func GetAllUser(offset, limit int) ([]model.CoreUser, int64, error) {
	db := config.DB
	var count int64
	var users []model.CoreUser
	result := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&users)
	db.Model(&model.CoreUser{}).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return users, count, nil
}

func GetUserById(id string) (*model.CoreUser, error) {
	db := config.DB
	var user model.CoreUser
	if err := db.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(e string) (*model.CoreUser, error) {
	db := config.DB
	var user model.CoreUser
	if err := db.Where(&model.CoreUser{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(u string) (*model.CoreUser, error) {
	db := config.DB
	var user model.CoreUser
	if err := db.Where(&model.CoreUser{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(data *model.CoreUser) (*model.CoreUser, error) {
	db := config.DB
	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func UpdateUser(user *model.CoreUser, payload *model.CoreUser) (*model.CoreUser, error) {
	db := config.DB
	updates := make(map[string]interface{})
	if payload.Username != "" {
		updates["username"] = payload.Username
	}
	if payload.Email != "" {
		updates["password"] = payload.Email
	}
	if err := db.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DestroyUser(id string) (result *gorm.DB) {
	db := config.DB
	result = db.Unscoped().Delete(&model.CoreUser{}, "id = ?", id)
	return result
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func ValidToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))
	if uid != n {
		return false
	}
	return true
}
