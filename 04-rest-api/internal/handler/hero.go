package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/config"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/entity"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/validator"
)

type HeroHandler struct {
	app *config.Application
}

func (h *HeroHandler) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	heroes, err := h.app.Entity.Hero.GetAll()
	if err != nil {
		h.app.Error.InternalServer(w, err)
		return
	}

	data := map[string]any{"data": heroes}
	if err := h.app.Util.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *HeroHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var input struct {
		Name      string `json:"name"`
		Universe  string `json:"universe"`
		Image_url string `json:"image_url"`
	}

	// read body
	err := h.app.Util.ReadJSON(w, r, &input)
	if err != nil {
		h.app.Error.BadRequest(w, r, err)
		return
	}

	hero := entity.Hero{
		Name:     input.Name,
		Universe: input.Universe,
		ImageURL: input.Image_url,
	}

	// validate
	v := validator.New()
	if entity.ValidateHero(v, &hero); !v.Valid() {
		h.app.Error.FailedValidation(w, v.Errors)
		return
	}

	// insert
	hero.HeroID, err = h.app.Entity.Hero.Insert(&hero)
	if err != nil {
		h.app.Error.InternalServer(w, err)
	}

	data := map[string]any{"data": hero}
	if err := h.app.Util.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
