//go:build wireinject
// +build wireinject

package app

import (
	"github.com/cch123/go-ddd/domain"
	"github.com/cch123/go-ddd/infra/config/static"
	"github.com/cch123/go-ddd/repo"
	"github.com/google/wire"
)

func initApp() (*app, error) {
	panic(wire.Build(newApp, domain.ProviderSet, repo.ProviderSet, static.ReadConfig))
}
