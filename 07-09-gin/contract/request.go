package contract

import (
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/model"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/validator"
)

/* ---------------------------------------------------------------- */
/*                               login                              */
/* ---------------------------------------------------------------- */

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r ReqLogin) Validate() (errors map[string]string) {
	v := validator.New()
	v.Check(r.Email != "", "email", "must be provided")
	v.Check(validator.Matches(r.Email, validator.EmailRX), "email", "must be valid email form")
	v.Check(r.Password != "", "password", "must be provided")
	v.Check(len(r.Password) >= 8, "password", "must be min 8 character long")

	if !v.Valid() {
		return v.Errors
	}
	return nil
}

func (r ReqLogin) ToUser() *model.Store {
	user := &model.Store{
		Email: r.Email,
	}
	return user
}

/* ---------------------------------------------------------------- */
/*                             register                             */
/* ---------------------------------------------------------------- */

type ReqRegister struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Type     *string `json:"type"`
}

func (r ReqRegister) Validate() (errors map[string]string) {
	v := validator.New()

	v.Check(r.Email != "", "email", "must be provided")
	v.Check(validator.Matches(r.Email, validator.EmailRX), "email", "must be valid email form")

	v.Check(r.Password != "", "password", "must be provided")
	v.Check(len(r.Password) >= 8, "password", "must be min 8 character long")

	v.Check(r.Name != "", "name", "must be provided")
	v.Check(len(r.Name) >= 6, "name", "must be min 6 character long")
	v.Check(len(r.Name) <= 15, "name", "must be max 15 character long")

	if r.Type != nil {
		v.Check(
			validator.PermittedValue[string](
				*r.Type,
				string(model.Silver),
				string(model.Platinum),
				string(model.Gold)),
			"type",
			"must be either silver, gold or platinum",
		)
	}

	if !v.Valid() {
		return v.Errors
	}
	return nil
}

func (r ReqRegister) ToUser() *model.Store {
	user := &model.Store{
		Email: r.Email,
		Name:  r.Name,
	}
	user.Password.Set(r.Password)

	if r.Type == nil {
		user.Type = model.Silver
	} else {
		user.Type = model.StoreType(*r.Type)
	}

	return user
}

/* ---------------------------------------------------------------- */
/*                          create product                          */
/* ---------------------------------------------------------------- */
type ReqPostProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Price       int    `json:"price"`
}

func (r ReqPostProduct) Validate() (errors map[string]string) {
	v := validator.New()
	v.Check(r.Name != "", "name", "must be provided")

	v.Check(r.Description != "", "description", "must be provided")
	v.Check(len(r.Description) >= 10, "description", "must be min 10 character long")

	v.Check(r.ImageUrl != "", "image_url", "must be provided")
	v.Check(r.Price >= 0, "price", "must be provided and positive greater than 0")

	if !v.Valid() {
		return v.Errors
	}
	return nil
}

func (r ReqPostProduct) ToProduct() *model.Product {
	product := &model.Product{
		Name:        r.Name,
		Description: r.Description,
		ImageUrl:    r.ImageUrl,
		Price:       r.Price,
	}
	return product
}

func (r ReqPostProduct) ToProductFromExisting(oldproduct *model.Product) *model.Product {
	oldproduct.Name = r.Name
	oldproduct.Description = r.Description
	oldproduct.ImageUrl = r.ImageUrl
	oldproduct.Price = r.Price
	return oldproduct
}
