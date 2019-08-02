package portfolio

import "github.com/FernandoCagale/c4-portfolio/pkg/entity"

//UseCase use case interface
type UseCase interface {
	FindAll(query *entity.Query) (values []*entity.Portfolio, count int, err error)
}
