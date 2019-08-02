package portfolio

import (
	"github.com/FernandoCagale/c4-portfolio/config"
	"github.com/FernandoCagale/c4-portfolio/pkg/entity"
)

//GormRepository in memory repo
type GormRepository struct {
	env *config.Config
}

//NewGormRepository create new repository
func NewGormRepository(env *config.Config) *GormRepository {
	return &GormRepository{env}
}

//FindAll all
func (r *GormRepository) FindAll(query *entity.Query) (values []*entity.Portfolio, count int, err error) {

	values = []*entity.Portfolio{
		&entity.Portfolio{
			Name: "Jhon",
			Actives: []*entity.Active{
				&entity.Active{
					Code: "FB",
					Operations: []*entity.Operation{
						&entity.Operation{
							Quantity: 10,
							Value:    1000,
						},
						&entity.Operation{
							Quantity: 12,
							Value:    1400,
						},
					},
				},
				&entity.Active{
					Code: "AMZN",
					Operations: []*entity.Operation{
						&entity.Operation{
							Quantity: 15,
							Value:    300,
						},
						&entity.Operation{
							Quantity: 12,
							Value:    200,
						},
					},
				},
			},
		},
		&entity.Portfolio{
			Name: "Max",
			Actives: []*entity.Active{
				&entity.Active{
					Code: "AMZN",
					Operations: []*entity.Operation{
						&entity.Operation{
							Quantity: 30,
							Value:    500,
						},
					},
				},
			},
		},
	}

	return values, 10, nil
}
