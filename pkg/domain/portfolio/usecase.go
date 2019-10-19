package portfolio

import "github.com/FernandoCagale/c4-portfolio/pkg/entity"

type UseCase interface {
	Create(task *entity.Portfolio) (err error)
	Update(id int, portfolio *entity.Portfolio) (err error)
	Delete(id int) (err error)
	FindByID(id int) (portfolio *entity.Portfolio, err error)
	FindAll() (portfolios []*entity.Portfolio, err error)
}
