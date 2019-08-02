package response

import "github.com/FernandoCagale/c4-portfolio/pkg/entity"

//FindAllResponse dto
type FindAllResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

//Data options
type Data struct {
	Count int                 `json:"count"`
	Rows  []*entity.Portfolio `json:"rows"`
}

//NewResponseFinAll findAll
func NewResponseFinAll(count int, rows []*entity.Portfolio) *Data {
	return &Data{count, convertResponseFinAllCanal(rows)}
}

func convertResponseFinAllCanal(portfolios []*entity.Portfolio) []*entity.Portfolio {
	//TODO
	return portfolios
}
