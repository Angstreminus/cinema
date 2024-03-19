package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Angstreminus/cinema/config"
	"github.com/Angstreminus/cinema/internal/apperrors"
	"github.com/Angstreminus/cinema/internal/auth"
	"github.com/Angstreminus/cinema/internal/dto"
	"github.com/Angstreminus/cinema/internal/service"
	"github.com/Angstreminus/cinema/logger"
)

type UserHandler struct {
	UserService *service.UserService
	Logger      *logger.Logger
	Config      *config.Config
}

func NewUserHandler(usrServ *service.UserService, log *logger.Logger, cfg *config.Config) *UserHandler {
	return &UserHandler{
		UserService: usrServ,
	}
}

func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	toRegistrate := &dto.RegisterRequest{}

	if err := json.NewDecoder(r.Body).Decode(toRegistrate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := uh.UserService.RegiterUser(toRegistrate)
	if err != nil {
		respErr := apperrors.MatchError(err)
		w.WriteHeader(respErr.Status)
		_ = json.NewEncoder(w).Encode(respErr)
		return
	}

	token, err := auth.CreateToken(uh.Config, user)
	if err != nil {
		respErr := apperrors.MatchError(err)
		w.WriteHeader(respErr.Status)
		_ = json.NewEncoder(w).Encode(respErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)

	if err = json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
}

func (uh *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	toLogin := &dto.LoginRequest{}

	if err := json.NewDecoder(r.Body).Decode(toLogin); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
