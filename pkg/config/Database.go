package config

import (
	"example/pkg/model"
	"fmt"
	//"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	//port, _ := strconv.ParseUint(GetEnv("DB_PORT"), 10, 32)
	//dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", GetEnv("DB_HOST"), port, GetEnv("DB_USER"), GetEnv("DB_PASSWORD"), GetEnv("DB_NAME"))
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", GetEnv("DB_USER"), GetEnv("DB_PASSWORD"), GetEnv("DB_HOST"), GetEnv("DB_PORT"), GetEnv("DB_NAME"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	fmt.Println("Connection success")
	err = DB.AutoMigrate(
		&model.CoreUser{},
		&model.InvSupplier{},
		&model.InvProductCategory{},
		&model.InvProduct{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}
