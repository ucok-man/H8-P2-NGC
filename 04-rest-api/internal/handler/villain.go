package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/config"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/entity"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/validator"
)

type VillainHandler struct {
	app *config.Application
}

func (h *VillainHandler) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	villains, err := h.app.Entity.Villain.GetAll()
	if err != nil {
		h.app.Error.InternalServer(w, err)
		return
	}

	data := map[string]any{"data": villains}
	if err := h.app.Util.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *VillainHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	villain := entity.Villain{
		Name:     input.Name,
		Universe: input.Universe,
		ImageURL: input.Image_url,
	}

	// validate
	v := validator.New()
	if entity.ValidateVillain(v, &villain); !v.Valid() {
		h.app.Error.FailedValidation(w, v.Errors)
		return
	}

	// insert
	villain.VillainID, err = h.app.Entity.Villain.Insert(&villain)
	if err != nil {
		h.app.Error.InternalServer(w, err)
	}

	data := map[string]any{"data": villain}
	if err := h.app.Util.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
