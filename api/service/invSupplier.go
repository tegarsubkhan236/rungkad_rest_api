package service

import (
	"example/pkg/config"
	"example/pkg/model"
	"gorm.io/gorm"
)

func GetAllSupplier(offset, limit int) ([]model.InvSupplier, int64, error) {
	var db = config.DB
	var count int64
	var data []model.InvSupplier

	result := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&data)
	db.Model(&model.InvSupplier{}).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return data, count, nil
}

func GetSupplierById(id string) (*model.InvSupplier, error) {
	var db = config.DB
	var item model.InvSupplier

	if err := db.Find(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func GetSupplierByColumn(data *model.InvSupplier, offset, limit int) ([]model.InvSupplier, int64, error) {
	var db = config.DB
	var modelData []model.InvSupplier
	var count int64

	if data.Name != "" {
		db = db.Where("name LIKE ?", "%"+data.Name+"%")
	}
	if err := db.Limit(limit).Offset(offset).Find(&modelData).Error; err != nil {
		return nil, 0, err
	}
	db.Model(&model.InvSupplier{}).Count(&count)
	return modelData, count, nil
}

func CreateSupplier(data *model.InvSupplier) (*model.InvSupplier, error) {
	var db = config.DB

	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func UpdateSupplier(item *model.InvSupplier, payload *model.InvSupplier) (*model.InvSupplier, error) {
	var db = config.DB

	if err := db.Model(&item).Updates(payload).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func DestroySupplier(id string) (result *gorm.DB) {
	var db = config.DB

	result = db.Unscoped().Delete(&model.InvSupplier{}, "id = ?", id)
	return result
}
