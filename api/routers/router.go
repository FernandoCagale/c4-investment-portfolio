package routers

import (
	"github.com/FernandoCagale/c4-portfolio/api/handlers"
	"github.com/gorilla/mux"
)

type SystemRoutes struct {
	healthHandler *handlers.HealthHandler
	portfolioHandler *handlers.PortfolioHandler
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/v1/api/portfolio/{id}", routes.portfolioHandler.FindByID).Methods("GET")
	r.HandleFunc("/v1/api/portfolio/{id}", routes.portfolioHandler.UpdateByID).Methods("PUT")
	r.HandleFunc("/v1/api/portfolio/{id}", routes.portfolioHandler.DeleteByID).Methods("DELETE")
	r.HandleFunc("/v1/api/portfolio", routes.portfolioHandler.FindAll).Methods("GET")
	r.HandleFunc("/v1/api/portfolio", routes.portfolioHandler.Create).Methods("POST")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, portfolioHandler *handlers.PortfolioHandler) *SystemRoutes {
	return &SystemRoutes{
		healthHandler: healthHandler,
		portfolioHandler: portfolioHandler,
	}
}
