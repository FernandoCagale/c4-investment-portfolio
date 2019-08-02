package routers

import (
	"github.com/FernandoCagale/c4-portfolio/config"
	"github.com/FernandoCagale/c4-portfolio/pkg/portfolio"
	"github.com/FernandoCagale/c4-portfolio/pkg/infra/infra"
	"github.com/gorilla/mux"
)

//NewRouter infra
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = infra.NotFoundHandler()
	router.MethodNotAllowedHandler = infra.NotAllowedHandler()
	return router
}

//makeGorm database postgres
func makeGorm(env *config.Config) portfolio.UseCase {
	return portfolio.NewService(portfolio.NewGormRepository(env))
}
