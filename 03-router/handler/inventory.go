package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ucok-man/H8-P2-NGC/03-router/config"
	"github.com/ucok-man/H8-P2-NGC/03-router/entity"
)

func GetInventories(app *config.Application) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		inventories, err := app.Entity.Inventory.GetAll()
		if err != nil {
			errorServer(w, err)
			return
		}

		err = writeJSON(w, http.StatusOK, map[string]any{
			"message": "Success!",
			"data":    inventories,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func GetInventoryByID(app *config.Application) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil || id < 1 {
			errorClient(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}

		inventory, err := app.Entity.Inventory.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				errorClient(w, http.StatusNotFound, http.StatusText(http.StatusBadRequest))
			default:
				errorServer(w, err)
			}
			return
		}

		err = writeJSON(w, http.StatusOK, map[string]any{
			"message": "Success!",
			"data":    inventory,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}

func CreateInventory(app *config.Application) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var inventory entity.Inventory
		err := readJSON(w, r, &inventory)
		if err != nil {
			errorClient(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}

		if inventory.Stock < 0 {
			errorClient(w, http.StatusBadRequest, "stock must be postive integer")
			return
		}

		if inventory.Status != "active" && inventory.Status != "broken" {
			errorClient(w, http.StatusBadRequest, "status must be active or broken")
			return
		}

		id, err := app.Entity.Inventory.Insert(&inventory)
		if err != nil {
			errorServer(w, err)
			return
		}
		inventory.ID = id

		err = writeJSON(w, http.StatusOK, map[string]any{
			"message": "Success!",
			"data":    inventory,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func UpdateInventory(app *config.Application) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var newInventory entity.Inventory

		id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil || id < 1 {
			errorClient(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		newInventory.ID = id

		err = readJSON(w, r, &newInventory)
		if err != nil {
			errorClient(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}

		if newInventory.Stock < 0 {
			errorClient(w, http.StatusBadRequest, "stock must be postive integer")
			return
		}

		if newInventory.Status != "active" && newInventory.Status != "broken" {
			errorClient(w, http.StatusBadRequest, "status must be active or broken")
			return
		}

		err = app.Entity.Inventory.UpdateByID(&newInventory)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				errorClient(w, http.StatusNotFound, http.StatusText(http.StatusBadRequest))
			default:
				errorServer(w, err)
			}
			return
		}

		err = writeJSON(w, http.StatusOK, map[string]any{
			"message": "Success!",
			"data":    newInventory,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func DeleteInventory(app *config.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
		if err != nil || id < 1 {
			errorClient(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}

		err = app.Entity.Inventory.DeleteByID(id)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				errorClient(w, http.StatusNotFound, http.StatusText(http.StatusBadRequest))
			default:
				errorServer(w, err)
			}
			return
		}

		err = writeJSON(w, http.StatusOK, map[string]any{
			"message": "Success!",
			"data":    fmt.Sprintf("inventory with id %d is deleted", id),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
