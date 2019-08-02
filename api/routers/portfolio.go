package routers

import (
	"github.com/FernandoCagale/c4-portfolio/api/handlers/portfolio"
	"github.com/FernandoCagale/c4-portfolio/config"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//MakeHandlers make url handlers
func MakeHandlers(r *mux.Router, n negroni.Negroni, env *config.Config) {
	service := makeGorm(env)

	r.Handle("/api/v1/portfolio", n.With(
		negroni.Wrap(portfolio.FindAll(service)),
	)).Methods("GET")

	r.Handle("/api/health", n.With(
		negroni.Wrap(portfolio.Health(service)),
	)).Methods("GET")
}
