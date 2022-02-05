package model

import "time"

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	ProductID int     `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	UserID    int     `json:"user_id"`
	User      User    `gorm:"foreingKey:UserID"`
}
