package repository

import (
	"github.com/CracherX/favorite_hist/internal/entity"
	"gorm.io/gorm"
)

type FavoriteRepository struct {
	DB *gorm.DB
}

func NewFavoriteRepoGorm(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{DB: db}
}

// GetFavoritesByUserID возвращает список всех избранных товаров для пользователя по его ID
func (repo *FavoriteRepository) GetFavoritesByUserID(userID int) ([]entity.Favorite, error) {
	var favorites []entity.Favorite

	// Выполняем запрос к БД, фильтруем по UserID
	if err := repo.DB.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		return nil, err
	}

	return favorites, nil
}

func (repo *FavoriteRepository) DeleteFavorite(userID, favID int) error {
	var favorite entity.Favorite

	result := repo.DB.Where("user_id = ? AND product_id = ?", userID, favID).Delete(&favorite)
	if result.RowsAffected == 0 {
		return ErrRecordNotFound // Ошибка, если запись не найдена
	}
	if result.Error != nil {
		return result.Error // Возвращаем ошибку, если что-то пошло не так
	}
	return nil // Успешное удаление
}

func (repo *FavoriteRepository) AddFavorite(userID, prodID int) error {
	favorite := entity.Favorite{
		ProductID: prodID,
		UserID:    userID,
	}

	result := repo.DB.Create(&favorite)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
