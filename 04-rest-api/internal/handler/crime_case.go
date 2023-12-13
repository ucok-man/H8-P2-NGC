package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/config"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/entity"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/validator"
)

type CrimeCaseHandler struct {
	app *config.Application
}

func (h *CrimeCaseHandler) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	crimecases, err := h.app.Entity.CrimeCase.GetAll()
	if err != nil {
		h.app.Error.InternalServer(w, err)
		return
	}

	for _, crimecase := range crimecases {
		hero, err := h.app.Entity.Hero.GetByID(crimecase.Hero.HeroID)
		if err != nil {
			h.app.Error.InternalServer(w, err)
			return
		}
		crimecase.Hero = *hero

		villain, err := h.app.Entity.Villain.GetByID(crimecase.Villain.VillainID)
		if err != nil {
			h.app.Error.InternalServer(w, err)
			return
		}
		crimecase.Villain = *villain
	}

	data := map[string]any{"data": crimecases}
	if err := h.app.Util.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *CrimeCaseHandler) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	crimecaseId, err := paramID(w, p)
	if err != nil {
		h.app.Error.BadRequest(w, r, err)
		return
	}

	crimeCase, err := h.app.Entity.CrimeCase.GetByID(crimecaseId)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.app.Error.NotFound(w, r)
		default:
			h.app.Error.InternalServer(w, err)
		}
		return
	}

	hero, err := h.app.Entity.Hero.GetByID(crimeCase.Hero.HeroID)
	if err != nil {
		h.app.Error.InternalServer(w, err)
		return
	}
	crimeCase.Hero = *hero

	villain, err := h.app.Entity.Villain.GetByID(crimeCase.Villain.VillainID)
	if err != nil {
		h.app.Error.InternalServer(w, err)
		return
	}
	crimeCase.Villain = *villain

	data := map[string]any{"data": crimeCase}
	if err := h.app.Util.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *CrimeCaseHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var input struct {
		HeroID       int64     `json:"hero_id"`
		VillainID    int64     `json:"villain_id"`
		Description  string    `json:"description"`
		IncidentDate time.Time `json:"incident_date"`
	}

	// read body
	err := h.app.Util.ReadJSON(w, r, &input)
	if err != nil {
		h.app.Error.BadRequest(w, r, err)
		return
	}

	crimeCase := entity.CrimeCase{
		Hero:         entity.Hero{HeroID: input.HeroID},
		Villain:      entity.Villain{VillainID: input.VillainID},
		Description:  input.Description,
		IncidentDate: input.IncidentDate,
	}

	// validate
	v := validator.New()
	if entity.ValidateCrimeCase(v, &crimeCase); !v.Valid() {
		h.app.Error.FailedValidation(w, v.Errors)
		return
	}

	_, err = h.app.Entity.Hero.GetByID(crimeCase.Hero.HeroID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.app.Error.NotFound(w, r)
		default:
			h.app.Error.InternalServer(w, err)
		}
		return
	}

	_, err = h.app.Entity.Villain.GetByID(crimeCase.Villain.VillainID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.app.Error.NotFound(w, r)
		default:
			h.app.Error.InternalServer(w, err)
		}
		return
	}

	// insert
	crimeCase.CrimeCaseID, err = h.app.Entity.CrimeCase.Insert(&crimeCase)
	if err != nil {
		h.app.Error.InternalServer(w, err)
	}

	// send
	response := struct {
		CrimeCaseID  int64     `json:"crime_case_id"`
		HeroID       int64     `json:"hero_id"`
		VillainID    int64     `json:"villain_id"`
		Description  string    `json:"description"`
		IncidentDate time.Time `json:"incident_date"`
	}{
		CrimeCaseID:  crimeCase.CrimeCaseID,
		HeroID:       crimeCase.Hero.HeroID,
		VillainID:    crimeCase.Villain.VillainID,
		Description:  crimeCase.Description,
		IncidentDate: crimeCase.IncidentDate,
	}

	data := map[string]any{"data": response}
	if err := h.app.Util.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *CrimeCaseHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	crimecaseId, err := paramID(w, p)
	if err != nil {
		h.app.Error.BadRequest(w, r, err)
		return
	}

	// get old record
	crimeCase, err := h.app.Entity.CrimeCase.GetByID(crimecaseId)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.app.Error.NotFound(w, r)
		default:
			h.app.Error.InternalServer(w, err)
		}
		return
	}

	// read input
	var input struct {
		HeroID       int64     `json:"hero_id"`
		VillainID    int64     `json:"villain_id"`
		Description  string    `json:"description"`
		IncidentDate time.Time `json:"incident_date"`
	}

	err = h.app.Util.ReadJSON(w, r, &input)
	if err != nil {
		h.app.Error.BadRequest(w, r, err)
		return
	}

	// mapping
	crimeCase.Hero.HeroID = input.HeroID
	crimeCase.Villain.VillainID = input.VillainID
	crimeCase.Description = input.Description
	crimeCase.IncidentDate = input.IncidentDate

	// validate
	v := validator.New()
	if entity.ValidateCrimeCase(v, crimeCase); !v.Valid() {
		h.app.Error.FailedValidation(w, v.Errors)
		return
	}

	// update
	err = h.app.Entity.CrimeCase.Update(crimeCase)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.app.Error.NotFound(w, r)
		default:
			h.app.Error.InternalServer(w, err)
		}
		return
	}

	// send
	data := map[string]any{"data": fmt.Sprintf("record with %d id is updated", crimecaseId)}
	if err := h.app.Util.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *CrimeCaseHandler) DeleteByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	crimecaseId, err := paramID(w, p)
	if err != nil {
		h.app.Error.BadRequest(w, r, err)
		return
	}

	err = h.app.Entity.CrimeCase.DeleteByID(crimecaseId)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.app.Error.NotFound(w, r)
		default:
			h.app.Error.InternalServer(w, err)
		}
		return
	}

	data := map[string]any{"data": fmt.Sprintf("record with %d id is deleted", crimecaseId)}
	if err := h.app.Util.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
