package main

import (
	"fmt"
	"os"

	"github.com/ucok-man/H8-P2-NGC/01-reflect/validator"
)

func main() {
	hero := struct {
		Name  string `required:"true" min:"1" max:"20"`
		Age   int    `required:"true" min:"18" max:"45"`
		Email string `required:"true" isEmail:"true"`
	}{
		Name:  "Hulk",
		Age:   45,
		Email: "hulk@hulk.com",
	}

	if err := validator.New(hero).Validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Validation OK!")
}
