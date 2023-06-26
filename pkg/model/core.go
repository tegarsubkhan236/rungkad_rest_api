package model

import (
	"gorm.io/gorm"
)

type CoreUser struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"type:varchar(32);unique_index;not null" json:"username"`
	Email    string `gorm:"type:varchar(32);unique_index;not null" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	gorm.Model
}
