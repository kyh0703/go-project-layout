//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	NewApp,
)

func InitializeApp(ctx context.Context) (*App, error) {
	panic(wire.Build(AppSet))
}
