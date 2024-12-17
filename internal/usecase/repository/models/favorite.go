package models

type Favorite struct {
	ID        int `gorm:"primaryKey;autoIncrement"`   // Первичный ключ с автоинкрементом
	ProductID int `gorm:"column:product_id;not null"` // Внешний ключ на продукт
	UserID    int `gorm:"column:user_id;not null"`    // Внешний ключ на пользователя
}
