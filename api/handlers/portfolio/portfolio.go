package portfolio

import (
	"net/http"

	"github.com/FernandoCagale/c4-portfolio/api/handlers/portfolio/response"
	"github.com/FernandoCagale/c4-portfolio/api/infra"
	"github.com/FernandoCagale/c4-portfolio/pkg/entity"
	"github.com/FernandoCagale/c4-portfolio/pkg/infra/errors"
	"github.com/FernandoCagale/c4-portfolio/pkg/infra/render"
	"github.com/FernandoCagale/c4-portfolio/pkg/portfolio"
)

//FindAll findAll
func FindAll(service portfolio.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")

		limit, err := infra.GetLimit(r)
		if err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Failed to decode request query"))
			return
		}

		page, err := infra.GetPage(r)
		if err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Failed to decode request query"))
			return
		}

		query := entity.NewQuery(search, limit, page)

		portfolios, count, err := service.FindAll(query)
		if err != nil {
			render.ResponseError(w, errors.Err(err))
			return
		}

		render.Response(w, response.NewResponseFinAll(count, portfolios), http.StatusOK)
	})
}

//Health health
func Health(service portfolio.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.Response(w, map[string]bool{"ok": true}, http.StatusOK)
	})
}
