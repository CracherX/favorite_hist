package dto

type GetUserFavoriteRequest struct {
	JWT string `validate:"required,jwt"`
}

type DeleteFavoriteRequest struct {
	JWT        string `validate:"required,jwt"`
	FavoriteId string `validate:"required,numeric"`
}

type AddFavoriteRequest struct {
	JWT       string `validate:"required,jwt"`
	ProductId string `validate:"required,numeric"`
}
