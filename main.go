package main

import (
	"log"
	"net/http"
	"github.com/Go-Run-React/Catalog-Microservice-Backend/handlers"
	m "github.com/Go-Run-React/Catalog-Microservice-Backend/middleware"
	u "github.com/Go-Run-React/Catalog-Microservice-Backend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var l *log.Logger

func dataRoutes(r chi.Router){
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", handlers.GetCategories)
		r.Post("/", handlers.AddCategory)
	})

	r.Route("/taxclassifications", func(r chi.Router) {
		r.Get("/", handlers.GetTaxClassifications)
		r.Post("/", handlers.AddTaxClassification)
	})
}

func main() {
	//General logger
	l = u.Glog

	//chi router
    r := chi.NewRouter()
	r.Use(middleware.Logger)

	//index routing
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to Catalog-Data-Backend"))
    })

	//new router for data-apis
	dataRouter := chi.NewRouter()
	dataRouter.Use(m.Repository)
	dataRoutes(dataRouter)
	
	r.Mount("/data", dataRouter)

	//server activation
    http.ListenAndServe(":3000", r)
}