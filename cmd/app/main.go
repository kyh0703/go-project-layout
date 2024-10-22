package main

import (
	"github.com/kyh0703/go-project-layout/configs"
	"github.com/kyh0703/go-project-layout/internal/app/presenter"
	"go.uber.org/fx"
)

// @title go-project-layout API
// @version 1.0
// @host localhost:8080
// @accept application/json
// @produce application/json
func main() {
	fx.New(
		configs.Module,
		presenter.Module,
	).Run()
}
