package portfolio

import "github.com/google/wire"

var Set = wire.NewSet(NewUseCase, NewRepository, wire.Bind(new(Repository), new(*PortfolioRepository)))