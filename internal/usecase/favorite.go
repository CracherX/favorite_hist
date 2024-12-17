package usecase

import "github.com/CracherX/favorite_hist/internal/entity"

type FavoriteUseCase struct {
	repo FavoriteRepo
}

func NewFavoriteUC(repo FavoriteRepo) *FavoriteUseCase {
	return &FavoriteUseCase{repo: repo}
}

func (uc *FavoriteUseCase) GetUserFavorite(id int) ([]entity.Favorite, error) {
	return uc.repo.GetFavoritesByUserID(id)
}

func (uc *FavoriteUseCase) DeleteFavorite(userID, favID int) error {
	return uc.repo.DeleteFavorite(userID, favID)
}

func (uc *FavoriteUseCase) AddFavorite(userID, prodID int) error {
	return uc.repo.AddFavorite(userID, prodID)
}
