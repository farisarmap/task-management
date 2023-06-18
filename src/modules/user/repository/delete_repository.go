package repository

import (
	"context"
	"database/sql"
)

func (repo *UserRepository) Delete(ctx context.Context, tx *sql.Tx, id string) bool {
	panic("not implemented") // TODO: Implement
}
