package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ucok-man/H8-P2-NGC/03-router/config"
	"github.com/ucok-man/H8-P2-NGC/03-router/handler"
)

func main() {
	app, cleanfn := config.NewApp()
	defer cleanfn()

	router := httprouter.New()
	router.GET("/inventories", handler.GetInventories(app))
	router.GET("/inventories/:id", handler.GetInventoryByID(app))
	router.POST("/inventories", handler.CreateInventory(app))
	router.PUT("/inventories/:id", handler.UpdateInventory(app))
	router.DELETE("/inventories/:id", handler.DeleteInventory(app))

	srv := &http.Server{
		Addr:    app.Addr,
		Handler: router,
	}

	log.Printf("runing server on %v...\n", app.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
