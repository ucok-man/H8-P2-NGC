package handler

import (
	"errors"
	"net/http"

	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/config"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/entity"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/validator"
)

type RecipeHandler struct {
	*config.Application
}

type CustomerHandler struct {
	*config.Application
}

func (h *RecipeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// get record
	recipes, err := h.Entity.Recipe.GetAll()
	if err != nil {
		h.JSON.Error.InternalServer(w, err)
		return
	}

	if nil == recipes {
		h.JSON.Success.OK(w, "success get recipes", "recipes", nil)
		return
	}

	// send
	h.JSON.Success.OK(w, "success get recipes", "recipes", recipes)
}

func (h *RecipeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// get param id
	recipeID, err := paramID(r)
	if err != nil {
		h.JSON.Error.BadRequest(w, r, err)
		return
	}

	// get record
	recipe, err := h.Entity.Recipe.GetByID(recipeID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.JSON.Error.NotFound(w, r)
		default:
			h.JSON.Error.InternalServer(w, err)
		}
		return
	}

	// send
	h.JSON.Success.OK(w, "succes get recipe", "recipe", recipe)
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name         string              `json:"name"`
		Description  string              `json:"description"`
		TimeRequired entity.TimeRequired `json:"time_required"`
		Rating       int                 `json:"rating"`
	}

	// read body
	err := h.JSON.Unmarshall(w, r, &input)
	if err != nil {
		h.JSON.Error.BadRequest(w, r, err)
		return
	}

	// mapping
	recipe := entity.Recipe{
		Name:         input.Name,
		Description:  input.Description,
		TimeRequired: input.TimeRequired,
		Rating:       input.Rating,
	}

	// validate
	v := validator.New()
	if entity.ValidateRecipe(v, &recipe); !v.Valid() {
		h.JSON.Error.FailedValidation(w, v.Errors)
		return
	}

	// insert
	recipe.RecipeID, err = h.Entity.Recipe.Insert(&recipe)
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

	// send
	h.JSON.Success.Created(w, "success create recipe", "recipe", recipe)
}

func (h *RecipeHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	// get param id
	recipeID, err := paramID(r)
	if err != nil {
		h.JSON.Error.BadRequest(w, r, err)
		return
	}

	// get old record
	recipe, err := h.Entity.Recipe.GetByID(recipeID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.JSON.Error.NotFound(w, r)
		default:
			h.JSON.Error.InternalServer(w, err)
		}
		return
	}

	// read input
	var input struct {
		Name         string              `json:"name"`
		Description  string              `json:"description"`
		TimeRequired entity.TimeRequired `json:"time_required"`
		Rating       int                 `json:"rating"`
	}

	err = h.JSON.Unmarshall(w, r, &input)
	if err != nil {
		h.JSON.Error.BadRequest(w, r, err)
		return
	}

	// mapping
	recipe.Name = input.Name
	recipe.Description = input.Description
	recipe.TimeRequired = input.TimeRequired
	recipe.Rating = input.Rating

	// validate
	v := validator.New()
	if entity.ValidateRecipe(v, recipe); !v.Valid() {
		h.JSON.Error.FailedValidation(w, v.Errors)
		return
	}

	// update
	err = h.Entity.Recipe.Update(recipe)
	if err != nil {
		h.JSON.Error.InternalServer(w, err)
		return
	}

	// send
	h.JSON.Success.OK(w, "success update recipe", "recipe", recipe)
}

func (h *RecipeHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	// get param id
	customerID, err := paramID(r)
	if err != nil {
		h.JSON.Error.BadRequest(w, r, err)
		return
	}

	// get old record
	recipe, err := h.Entity.Recipe.GetByID(customerID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrorNotFound):
			h.JSON.Error.NotFound(w, r)
		default:
			h.JSON.Error.InternalServer(w, err)
		}
		return
	}

	// delete
	err = h.Entity.Recipe.DeleteByID(recipe.RecipeID)
	if err != nil {
		h.JSON.Error.InternalServer(w, err)
		return
	}

	// send
	h.JSON.Success.OK(w, "success delete recipe", "recipe", recipe)
}
