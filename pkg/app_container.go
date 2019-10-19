package pkg

import (
	"github.com/FernandoCagale/c4-portfolio/api/handlers"
	"github.com/FernandoCagale/c4-portfolio/api/routers"
	"github.com/FernandoCagale/c4-portfolio/pkg/domain/portfolio"
	"github.com/google/wire"
)

var Container = wire.NewSet(portfolio.Set, handlers.Set, routers.Set)
