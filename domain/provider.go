package domain

import (
	"github.com/cch123/go-ddd/domain/aggregate"
	"github.com/google/wire"
)

// ProviderSet all new aggregate should be registered here
var ProviderSet = wire.NewSet(aggregate.NewCustomerAgg, aggregate.NewOrderAgg)
