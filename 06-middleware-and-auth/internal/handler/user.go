package handler

import (
	"errors"
	"net/http"

	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/config"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/entity"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/jwt"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/validator"
)

type UserHandler struct {
	*config.Application
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input entity.LoginUser

	if err := h.JSON.Unmarshall(w, r, &input); err != nil {
		h.JSON.Error.BadRequest(w, r, err)
		return
	}

	v := validator.New()
	if entity.ValidateUserLogin(v, &input); !v.Valid() {
		h.JSON.Error.FailedValidation(w, v.Errors)
		return
	}

	existingUser, err := h.Entity.User.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.JSON.Error.InvalidCredentials(w, r)
		default:
			h.JSON.Error.InternalServer(w, err)
		}
		return
	}

	if err := existingUser.Password.Matches(input.Password); err != nil {
		h.JSON.Error.InvalidCredentials(w, r)
		return
	}

	claim := jwt.NewJWTClaim(existingUser.UserID)
	token, err := jwt.GenerateToken(&claim, h.Config.JWTSecreet)
	if err != nil {
		h.JSON.Error.InternalServer(w, err)
		return
	}

	h.JSON.Success.OK(w, "login success", "token", token)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Occupation string `json:"occupation"`
		Role       string `json:"role"`
	}

	if err := h.JSON.Unmarshall(w, r, &input); err != nil {
		h.JSON.Error.BadRequest(w, r, err)
		return
	}

	user := &entity.User{
		Email:      input.Email,
		Name:       input.Name,
		Age:        input.Age,
		Occupation: input.Occupation,
		Role:       input.Role,
	}
	user.Password.SetAndHash(input.Password)

	v := validator.New()
	if entity.ValidateUserRegister(v, user); !v.Valid() {
		h.JSON.Error.FailedValidation(w, v.Errors)
		return
	}

	var err error
	user.UserID, err = h.Entity.User.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrDuplicateEntry):
			v.AddError("email", "already exists")
			h.JSON.Error.FailedValidation(w, v.Errors)
		default:
			h.JSON.Error.InternalServer(w, err)
		}
		return
	}

	claim := jwt.NewJWTClaim(user.UserID)
	token, err := jwt.GenerateToken(&claim, h.Config.JWTSecreet)
	if err != nil {
		h.JSON.Error.InternalServer(w, err)
		return
	}

	h.JSON.Success.OK(w, "register success", "token", token)
}
