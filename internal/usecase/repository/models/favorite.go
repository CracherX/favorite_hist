package models

type Favorite struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	ProductID int `gorm:"column:product_id;not null;uniqueIndex:user_product_unique"`
	UserID    int `gorm:"column:user_id;not null;uniqueIndex:user_product_unique"`
}
