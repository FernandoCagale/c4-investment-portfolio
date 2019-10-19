package portfolio

import (
	"github.com/FernandoCagale/c4-portfolio/pkg/entity"
	"github.com/FernandoCagale/c4-portfolio/pkg/infra/errors"
)

type PortfolioUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *PortfolioUseCase {
	return &PortfolioUseCase{
		repo: repo,
	}
}

func (usecase *PortfolioUseCase) Create(e *entity.Portfolio) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	return usecase.repo.Create(e)
}

func (usecase *PortfolioUseCase) Update(id int, e *entity.Portfolio) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}
	
	return usecase.repo.Update(id, e)
}

func (usecase *PortfolioUseCase) Delete(id int) error {
	return usecase.repo.Delete(id)
}

func (usecase *PortfolioUseCase) FindByID(id int) (task *entity.Portfolio, err error) {
	return usecase.repo.FindByID(id)
}

func (usecase *PortfolioUseCase) FindAll() (tasks []*entity.Portfolio, err error) {
	return usecase.repo.FindAll()
}
