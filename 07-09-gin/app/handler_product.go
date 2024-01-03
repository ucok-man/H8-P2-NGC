package app

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/contract"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/repo"
)

func (app *Application) getAllProductHandler(ctx *gin.Context) {
	user := app.getCurrentUser(ctx)

	products, err := app.Repo.Product.GetAll(user.StoreID)
	if err != nil {
		contract.ErrInternalServer(ctx, err)
		return
	}

	if len(products) < 1 {
		contract.StatusOK(ctx, "currently we have no data yet", nil)
		return
	}
	contract.StatusOK(ctx, "OK", products)
}

func (app *Application) getProductByIdHandler(ctx *gin.Context) {
	user := app.getCurrentUser(ctx)

	productID, err := app.getParamId(ctx)
	if err != nil {
		contract.ErrBadRequest(ctx, err)
		return
	}

	product, err := app.Repo.Product.GetByID(user.StoreID, productID)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrRecordNotFound):
			contract.ErrNotFound(ctx)
		default:
			contract.ErrInternalServer(ctx, err)
		}
		return
	}

	contract.StatusOK(ctx, "OK", product)
}

func (app *Application) createProductHandler(ctx *gin.Context) {
	user := app.getCurrentUser(ctx)

	var input contract.ReqPostProduct

	if err := ctx.ShouldBindJSON(&input); err != nil {
		contract.ErrBadRequest(ctx, err)
		return
	}

	if err := input.Validate(); err != nil {
		contract.ErrFailedValidation(ctx, err)
		return
	}

	product := input.ToProduct()
	if err := app.Repo.Product.Insert(user.StoreID, product); err != nil {
		contract.ErrInternalServer(ctx, err)
		return
	}

	contract.StatusCreated(ctx, "Created", product)
}

func (app *Application) updateProductHandler(ctx *gin.Context) {
	user := app.getCurrentUser(ctx)

	productID, err := app.getParamId(ctx)
	if err != nil {
		contract.ErrBadRequest(ctx, err)
		return
	}

	product, err := app.Repo.Product.GetByID(user.StoreID, productID)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrRecordNotFound):
			contract.ErrNotFound(ctx)
		default:
			contract.ErrInternalServer(ctx, err)
		}
		return
	}

	var input contract.ReqPostProduct

	if err := ctx.ShouldBindJSON(&input); err != nil {
		contract.ErrBadRequest(ctx, err)
		return
	}

	if err := input.Validate(); err != nil {
		contract.ErrFailedValidation(ctx, err)
		return
	}

	product = input.ToProductFromExisting(product)

	if err := app.Repo.Product.Update(user.StoreID, product); err != nil {
		contract.ErrInternalServer(ctx, err)
		return
	}

	contract.StatusOK(ctx, "OK", product)
}

func (app *Application) deleteProductHandler(ctx *gin.Context) {
	user := app.getCurrentUser(ctx)

	productID, err := app.getParamId(ctx)
	if err != nil {
		contract.ErrBadRequest(ctx, err)
		return
	}

	product, err := app.Repo.Product.GetByID(user.StoreID, productID)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrRecordNotFound):
			contract.ErrNotFound(ctx)
		default:
			contract.ErrInternalServer(ctx, err)
		}
		return
	}

	err = app.Repo.Product.Delete(user.StoreID, product.StoreID)
	if err != nil {
		contract.ErrInternalServer(ctx, err)
		return
	}

	contract.StatusOK(ctx, "OK", product)
}
