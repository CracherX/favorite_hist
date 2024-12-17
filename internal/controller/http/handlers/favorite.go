package handlers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/CracherX/favorite_hist/internal/controller/http/dto"
	"github.com/CracherX/favorite_hist/internal/usecase/repository"
	"net/http"
	"strconv"
)

type FavoriteHandler struct {
	uc  FavoriteUC
	log Logger
	val Validator
	cl  Client
}

func NewFavoriteHandler(uc FavoriteUC, log Logger, val Validator, cl Client) *FavoriteHandler {
	return &FavoriteHandler{uc: uc, log: log, val: val, cl: cl}
}

func (h *FavoriteHandler) GetUserFavorite(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	data := dto.GetUserFavoriteRequest{JWT: query.Get("jwt")}

	if err := h.val.Validate(&data); err != nil {
		h.log.Debug("Bad Request", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	id, _ := h.auth(data.JWT, w)

	favorites, err := h.uc.GetUserFavorite(id)
	if err != nil {
		if errors.Is(err, driver.ErrBadConn) {
			h.log.Error("Bad Gateway", "Ошибка", err.Error())
			dto.Response(w, http.StatusBadGateway, "Ошибка в работе внешних сервисов", "Обратитесь к техническому специалисту")
		} else {
			h.log.Error("Bad Gateway", "Ошибка", err.Error())
			dto.Response(w, http.StatusInternalServerError, "Ошибка в работе сервера", "Обратитесь к техническому специалисту")
		}
		return
	}

	res := dto.UserFavoriteResponse{Favorites: favorites}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		h.log.Error("Ошибка энкодера", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Ошибка в работе сервера", "Обратитесь к техническому специалисту")
	}
}

func (h *FavoriteHandler) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	data := dto.DeleteFavoriteRequest{
		JWT:        query.Get("jwt"),
		FavoriteId: query.Get("id"),
	}

	if err := h.val.Validate(&data); err != nil {
		h.log.Debug("Bad Request", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	id, _ := strconv.Atoi(data.FavoriteId)

	userID, _ := h.auth(data.JWT, w)

	if err := h.uc.DeleteFavorite(userID, id); err != nil {
		switch {
		case errors.Is(err, driver.ErrBadConn):
			h.log.Error("Bad Gateway", "Ошибка", err.Error())
			dto.Response(w, http.StatusBadGateway, "Ошибка в работе внешних сервисов", "Обратитесь к техническому специалисту")
		case errors.Is(err, repository.ErrRecordNotFound):
			h.log.Debug("Bad Request", "Ошибка", err.Error())
			dto.Response(w, http.StatusNotFound, "Запись не найдена", "Убедитесь, что ID вы отправили именно ID избранного элемента (а не самого товара) и он точно принадлежит пользователю указанному в запросе")
		default:
			h.log.Error("Internal Server Error", "Ошибка", err.Error())
			dto.Response(w, http.StatusInternalServerError, "Внутренняя ошибка сервера", "Обратитесь к техническому специалисту")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")

	dto.Response(w, http.StatusOK, "Успешное удаление товара из избранного!")
}

func (h *FavoriteHandler) AddFavorite(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	data := dto.AddFavoriteRequest{
		JWT:       query.Get("jwt"),
		ProductId: query.Get("productID"),
	}

	if err := h.val.Validate(&data); err != nil {
		h.log.Debug("Bad Request", "Ошибка", err.Error())
		dto.Response(w, http.StatusBadRequest, "Bad Request", "Обратитесь к документации и заполните тело запроса правильно")
		return
	}

	pid, _ := strconv.Atoi(data.ProductId)

	uid, _ := h.auth(data.JWT, w)

	if err := h.uc.AddFavorite(uid, pid); err != nil {
		switch {
		case errors.Is(err, driver.ErrBadConn):
			h.log.Error("Bad Gateway", "Ошибка", err.Error())
			dto.Response(w, http.StatusBadGateway, "Ошибка в работе внешних сервисов", "Обратитесь к техническому специалисту")
		default:
			h.log.Debug("Duplicated Key попытка", "Ошибка", err.Error())
			dto.Response(w, http.StatusInternalServerError, "Внутренняя ошибка сервера", "Обратитесь к техническому специалисту")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")

	dto.Response(w, http.StatusOK, "Добавлено в избранное!")
}

func (h *FavoriteHandler) auth(jwt string, w http.ResponseWriter) (int, error) {
	var cdto dto.AuthClientResponse
	params := map[string]string{
		"jwt": jwt,
	}

	clr, err := h.cl.Get("/auth/profile", params)
	if err != nil {
		h.log.Error("Ошибка в работе клиента", "Запрос", "auth")
		dto.Response(w, http.StatusBadGateway, "Bad Gateway", "Проблема в работе внешних сервисов")
		return 0, err
	}

	if err = json.NewDecoder(clr.Body).Decode(&cdto); err != nil {
		h.log.Error("Ошибка работы энкодера", "Запрос", "auth", "Ошибка", err.Error())
		dto.Response(w, http.StatusInternalServerError, "Internal Server Error", "Внутренняя ошибка сервера, обратитесь к техническому специалисту")
		return 0, err
	}

	return cdto.UserID, nil
}
