package portfolio

import "github.com/FernandoCagale/c4-portfolio/pkg/entity"

//Repository repository interface
type Repository interface {
	FindAll(query *entity.Query) (values []*entity.Portfolio, count int, err error)
}
