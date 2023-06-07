package service

import (
	"example/pkg/config"
	"example/pkg/model"
)

func GetAllProduct(offset, limit int) ([]model.InvProduct, int64, error) {
	var db = config.DB
	var count int64
	var data []model.InvProduct

	result := db.Model(&model.InvProduct{}).Preload("InvSupplier").Preload("InvProductCategory").
		Count(&count).Limit(limit).Offset(offset).Find(&data)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return data, count, nil
}

func GetAllProductByFilter(offset, limit, supplierID int, productCategoryID []int, searchText string) ([]model.InvProduct, int64, error) {
	var db = config.DB
	var count int64
	var data []model.InvProduct

	db = db.Model(&model.InvProduct{})

	if supplierID != 0 {
		db = db.Preload("InvSupplier").Joins("JOIN inv_suppliers on inv_suppliers.id = inv_products.inv_supplier_id AND "+
			"inv_products.inv_supplier_id = ?", supplierID)
	} else {
		db = db.Preload("InvSupplier")
	}

	if len(productCategoryID) != 0 {
		db = db.Preload("InvProductCategory").Joins("JOIN inv_product_product_categories on inv_product_product_categories.inv_product_id = inv_products.id "+
			"JOIN inv_product_categories on inv_product_product_categories.inv_product_category_id = inv_product_categories.id "+
			"AND inv_product_categories.id in ? ", productCategoryID).Group("inv_products.id")
	} else {
		db = db.Preload("InvProductCategory")
	}

	if searchText != "" {
		db = db.Where("inv_products.name LIKE ?", "%"+searchText+"%")
	}

	result := db.Count(&count).Offset(offset).Limit(limit).Find(&data)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return data, count, nil
}

func GetProductById(id string) (*model.InvProduct, error) {
	var db = config.DB
	var item model.InvProduct

	if err := db.Preload("InvSupplier").Preload("InvProductCategory").Find(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func CreateProduct(data *model.InvProduct) (*model.InvProduct, error) {
	var db = config.DB

	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func UpdateProduct(item *model.InvProduct, payload *model.InvProduct) (*model.InvProduct, error) {
	var db = config.DB

	if err := db.Model(&item).Updates(payload).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func DestroyProduct(id string) error {
	var db = config.DB

	if err := db.Unscoped().Delete(&model.InvProduct{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
