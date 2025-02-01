package pkg

import (
	"github.com/kyh0703/layout/internal/pkg/db"
	"github.com/kyh0703/layout/internal/pkg/validate"
	"go.uber.org/fx"
)

var Module = fx.Module("pkg",
	fx.Provide(validate.NewValidator),
	fx.Provide(db.NewDB, db.NewQueries),
)
