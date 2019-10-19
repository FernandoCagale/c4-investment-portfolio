package handlers

import (
	"encoding/json"
	"github.com/FernandoCagale/c4-portfolio/api/render"
	"github.com/FernandoCagale/c4-portfolio/pkg/domain/portfolio"
	"github.com/FernandoCagale/c4-portfolio/pkg/entity"
	"github.com/FernandoCagale/c4-portfolio/pkg/infra/errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PortfolioHandler struct {
	usecase *portfolio.PortfolioUseCase
}

func NewPortfolio(usecase *portfolio.PortfolioUseCase) *PortfolioHandler {
	return &PortfolioHandler{
		usecase: usecase,
	}
}

func (handler *PortfolioHandler) Portfolios(w http.ResponseWriter, r *http.Request) {
	render.Response(w, map[string]bool{"ok": true}, http.StatusOK)
}

func (handler *PortfolioHandler) Create(w http.ResponseWriter, r *http.Request) {
	var portfolio *entity.Portfolio

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&portfolio); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := handler.usecase.Create(portfolio); err != nil {
		switch err {
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, portfolio, http.StatusCreated)
}

func (handler *PortfolioHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	var portfolio entity.Portfolio
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&portfolio); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err = handler.usecase.Update(id, &portfolio)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusNoContent)
}

func (handler *PortfolioHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	err = handler.usecase.Delete(id)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusNoContent)
}

func (handler *PortfolioHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	portfolio, err := handler.usecase.FindByID(id)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, portfolio, http.StatusOK)
}

func (handler *PortfolioHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	portfolios, err := handler.usecase.FindAll()
	if err != nil {
		render.ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	render.Response(w, portfolios, http.StatusOK)
}
