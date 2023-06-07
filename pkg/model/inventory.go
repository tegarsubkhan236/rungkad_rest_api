package model

import (
	"gorm.io/gorm"
	"time"
)

type InvSupplier struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"type:varchar(32);not null" json:"name"`
	Address       string         `gorm:"type:varchar(32);not null" json:"address"`
	ContactPerson string         `gorm:"type:varchar(32);not null" json:"contact_person"`
	ContactNumber string         `gorm:"type:varchar(32);not null" json:"contact_number"`
	Status        int8           `gorm:"default:0" json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type InvProductCategory struct {
	ID        uint                 `gorm:"primaryKey" json:"id"`
	Name      string               `gorm:"type:varchar(32);not null" json:"name"`
	ParentID  *uint                `json:"parent_id"`
	Children  []InvProductCategory `gorm:"foreignkey:ParentID;references:id" json:"children"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
}

type InvProduct struct {
	ID                 uint                 `gorm:"primaryKey" json:"id"`
	Name               string               `gorm:"type:varchar(32);not null" json:"name"`
	Image              string               `json:"image"`
	Cost               int                  `gorm:"not null" json:"cost"`
	Status             int8                 `gorm:"default:0" json:"status"`
	InvSupplierID      int                  `json:"inv_supplier_id"`
	InvSupplier        InvSupplier          `json:"inv_supplier"`
	Description        string               `gorm:"type:varchar(150)" json:"description"`
	InvProductCategory []InvProductCategory `gorm:"many2many:inv_product_product_categories" json:"inv_product_category"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	DeletedAt          gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
}

type InvStock struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	InvProductID int            `json:"inv_product_id"`
	InvProduct   InvProduct     `json:"inv_product"`
	CoreUserID   int            `json:"core_user_id"`
	CoreUser     CoreUser       `json:"core_user"`
	Quantity     int            `gorm:"default:0" json:"quantity"`
	Price        int            `json:"price"`
	Total        int            `json:"total"`
	Type         int            `json:"type"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type InvSell struct {
}

type InvPO struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type InvReceiving struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type InvBO struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type InvReturn struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
