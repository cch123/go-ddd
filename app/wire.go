//go:build wireinject
// +build wireinject

package app

import (
	"github.com/cch123/go-ddd/domain"
	"github.com/cch123/go-ddd/repo"
	"github.com/google/wire"
)

func initApp(i int) *app {
	panic(wire.Build(newApp, domain.ProviderSet, repo.ProviderSet))
}
