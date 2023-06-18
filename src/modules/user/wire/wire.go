//go:build wireinject
// +build wireinject

package wires

import (
	"database/sql"
	"task-management-be/src/modules/user/handler"

	"github.com/google/wire"
)

func ProvideUserInjection(db *sql.DB) *handler.UserController {
	wire.Build(ProviderSet)
	return nil
}
