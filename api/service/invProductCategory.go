package service

import (
	"errors"
	"example/pkg/config"
	"example/pkg/model"
	"gorm.io/gorm"
)

func GetAllProductCategory(offset, limit int) ([]model.InvProductCategory, int64, error) {
	db := config.DB
	var count int64
	var data []model.InvProductCategory
	result := db.Where("parent_id IS NULL").Limit(limit).Offset(offset).Preload("Children").Preload("Children.Children").Find(&data)
	if result.Error != nil {
		return nil, 0, errors.New("failed to get data")
	}
	db.Model(&model.InvProductCategory{}).Where("parent_id IS NULL").Count(&count)
	return data, count, nil
}

func GetProductCategoryById(id string) (*model.InvProductCategory, error) {
	db := config.DB
	var item model.InvProductCategory
	if err := db.Find(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func CreateProductCategory(data *model.InvProductCategory) (*model.InvProductCategory, error) {
	db := config.DB
	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func UpdateProductCategory(item *model.InvProductCategory, payload *model.InvProductCategory) (*model.InvProductCategory, error) {
	db := config.DB
	if err := db.Model(&item).Updates(payload).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func DestroyProductCategory(id string) (result *gorm.DB) {
	db := config.DB
	result = db.Unscoped().Delete(&model.InvProductCategory{}, "id = ?", id)
	return result
}
