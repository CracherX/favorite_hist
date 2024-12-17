package usecase

import "github.com/CracherX/favorite_hist/internal/entity"

type FavoriteRepo interface {
	GetFavoritesByUserID(userID int) ([]entity.Favorite, error)
	DeleteFavorite(userID, favID int) error
	AddFavorite(userID, prodID int) error
}
