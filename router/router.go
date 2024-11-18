package router

import (
	"log"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/train-do/Golang-Generics/database"
	"github.com/train-do/Golang-Generics/handler"
	"github.com/train-do/Golang-Generics/repository"
	"github.com/train-do/Golang-Generics/service"
)

func RouterAPI() *chi.Mux {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	repoDestination := repository.NewRepoDestination(db)
	serviceDestination := service.NewServiceDestination(repoDestination)
	handlerDestination := handler.NewRepoDestination(serviceDestination)
	repoOrder := repository.NewRepoOrder(db)
	serviceOrder := service.NewServiceOrder(repoOrder)
	handlerOrder := handler.NewRepoOrder(serviceOrder)
	repoLocation := repository.NewRepoLocation(db)
	serviceLocation := service.NewServiceLocation(repoLocation)
	handlerLocation := handler.NewRepoLocation(serviceLocation)
	repoTourPlan := repository.NewRepoTourPlan(db)
	serviceTourPlan := service.NewServiceTourPlan(repoTourPlan)
	handlerTourPlan := handler.NewRepoTourPlan(serviceTourPlan)
	router := chi.NewRouter()
	// router.Use(middleware.Logger)
	router.Get("/destinations", handlerDestination.GetAll)
	router.Get("/destination/{id}", handlerDestination.GetById)
	router.Post("/order", handlerOrder.AddOrder)
	router.Get("/location/{destinationId}", handlerLocation.GetByDestinationId)
	router.Get("/tour_plan/{destinationId}", handlerTourPlan.GetByDestinationId)
	// router.Route("/api", func(r chi.Router) {
	// 	r.Group(func(r chi.Router) {
	// 		r.Route("/cinemas", func(r chi.Router) {
	// 		})
	// 	})
	// })
	return router
}
