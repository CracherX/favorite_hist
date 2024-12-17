package entity

type Favorite struct {
	ID        int `json:"id"`        // Первичный ключ с автоинкрементом
	ProductID int `json:"productID"` // Внешний ключ на продукт
	UserID    int `json:"userID"`    // Внешний ключ на пользователя
}
