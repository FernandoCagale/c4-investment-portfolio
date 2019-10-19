package portfolio

import (
	"github.com/FernandoCagale/c4-portfolio/pkg/entity"
	"github.com/FernandoCagale/c4-portfolio/pkg/infra/errors"
	"strconv"
)

type PortfolioRepository struct {
	m map[string]*entity.Portfolio
}

func NewRepository() *PortfolioRepository {
	jhon := &entity.Portfolio{
		ID:   1,
		Name: "Jhon",
		Actives: []*entity.Active{
			{
				Code: "FB",
				Operations: []*entity.Operation{
					{
						Quantity: 10,
						Value:    1000,
					},
					{
						Quantity: 12,
						Value:    1400,
					},
				},
			},
			{
				Code: "AMZN",
				Operations: []*entity.Operation{
					{
						Quantity: 15,
						Value:    300,
					},
					{
						Quantity: 12,
						Value:    200,
					},
				},
			},
		},
	}

	max := &entity.Portfolio{
		ID:   2,
		Name: "Max",
		Actives: []*entity.Active{
			{
				Code: "AMZN",
				Operations: []*entity.Operation{
					{
						Quantity: 30,
						Value:    500,
					},
				},
			},
		},
	}

	var m = map[string]*entity.Portfolio{"1": jhon, "2": max}
	return &PortfolioRepository{m}
}

func (repo *PortfolioRepository) Create(e *entity.Portfolio) error {
	if repo.m[strconv.Itoa(e.ID)] != nil {
		return errors.ErrInvalidPayload
	}
	repo.m[strconv.Itoa(e.ID)] = e
	return nil
}

func (repo *PortfolioRepository) Update(id int, e *entity.Portfolio) error {
	if repo.m[strconv.Itoa(id)] == nil {
		return errors.ErrNotFound
	}
	repo.m[strconv.Itoa(id)] = e
	return nil
}

func (repo *PortfolioRepository) Delete(id int) error {
	if repo.m[strconv.Itoa(id)] == nil {
		return errors.ErrNotFound
	}
	delete(repo.m, strconv.Itoa(id))
	return nil
}

func (repo *PortfolioRepository) FindByID(id int) (task *entity.Portfolio, err error) {
	if repo.m[strconv.Itoa(id)] == nil {
		return nil, errors.ErrNotFound
	}
	return repo.m[strconv.Itoa(id)], nil
}

func (repo *PortfolioRepository) FindAll() (portfolios []*entity.Portfolio, err error) {
	for _, portfolio := range repo.m {
		portfolios = append(portfolios, portfolio)
	}
	return portfolios, nil
}
